package organizer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	cfg "github.com/fgeck/iptv-m3u-organizer/config"
	"github.com/fgeck/iptv-m3u-organizer/m3u"
)

type Organizer struct {
	logger     *log.Logger
	config     *cfg.Config
	parser     *m3u.M3uParser
	writer     *m3u.M3uWriter
	filter     *m3u.M3uFilter
	downloader *m3u.Downloader
}

func NewOrganizer(
	configFilePath string,
	url string,
	filter string,
	outputFilePath string,
	authType string,
) *Organizer {
	logger := log.New(os.Stdout, "organizer: ", log.Ldate|log.Ltime|log.Lshortfile)
	configReader := cfg.NewConfigReader()
	m3uParser := m3u.NewParser()
	m3uWriter := m3u.NewWriter()
	downloader := m3u.NewDownloader()
	m3uFilter := m3u.NewFilter()

	config := buildConfig(configReader, logger, configFilePath, url, filter, outputFilePath, authType)

	return &Organizer{
		logger,
		config,
		m3uParser,
		m3uWriter,
		m3uFilter,
		downloader,
	}
}

func buildConfig(
	configReader *cfg.ConfigReader,
	logger *log.Logger,
	configFilePath string,
	url string,
	filter string,
	outputFilePath string,
	authType string,
) *cfg.Config {
	var config *cfg.Config
	var err error

	if configFilePath != "" {
		logger.Printf("Using config file: %s\n", configFilePath)
		config, err = configReader.ConfigFromYaml(configFilePath)
		if err != nil {
			logger.Fatalf("failed to read yaml config: %v", err)
		}
	} else if url != "" && filter != "" && outputFilePath != "" && authType != "" {
		config, err = configReader.FilterFromString(filter)
		if err != nil {
			logger.Fatalf("failed to read yaml config: %v", err)
		}

		if authType != string(cfg.BasicAuth) && authType != string(cfg.URLParamAuth) {
			logger.Fatalf("Invalid auth type: %s", authType)
		}

		config.M3uURL = url
		config.OutputFilePath = outputFilePath
		logger.Printf("Parsed config from arguments: \n%v", config)
	} else {
		logger.Fatal("Invalid parameters. Either provide a config file or a URL, filter and outputFilePath.")
	}
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	if user == "" || password == "" {
		logger.Fatal("USER and PASSWORD environment variables must be set")
	}
	config.Auth = &cfg.AuthenticationInformation{
		User:     user,
		Password: password,
		AuthType: cfg.AuthType(authType),
	}

	return config
}

func (o *Organizer) Run() {
	m3uFilename := filepath.Join(m3u.DownloadsDir, fmt.Sprintf("%s.m3u", time.Now().Format("2006-01-02")))
	err := o.downloader.CreateDownloadsDirIfNotExist()
	if err != nil {
		o.logger.Fatalf("failed to create downloadDir: %v", err)
	}
	err = o.downloader.Download(o.config.M3uURL, m3uFilename, o.config.Auth)
	if err != nil {
		o.logger.Fatalf("failed to download m3u: %v", err)
	}

	m3u, err := o.parser.ParseM3U(m3uFilename)
	if err != nil {
		o.logger.Fatalf("failed to read m3u file: %v", err)
	}
	filteredM3u, err := o.filter.Filter(o.config, m3u)
	if err != nil {
		log.Fatalf("failed to filter m3u: %v", err)
	}

	if err = o.writer.WriteToFile(filteredM3u, o.config.OutputFilePath); err != nil {
		o.logger.Fatalf("failed to write to output file: %v", err)
	}
}
