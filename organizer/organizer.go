package organizer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fgeck/iptv-m3u-organizer/config"
	"github.com/fgeck/iptv-m3u-organizer/m3u"
)

type Organizer struct {
	logger     *log.Logger
	config     *config.Config
	parser     *m3u.M3uParser
	downloader *m3u.Downloader
}

func NewOrganizer(configFilePath string, url string, filter string) *Organizer {
	logger := log.New(os.Stdout, "organizer: ", log.Ldate|log.Ltime|log.Lshortfile)
	configReader := &config.ConfigReader{}
	m3uParser := &m3u.M3uParser{}
	downloader := m3u.NewDownloader()

	var config *config.Config
	var err error

	if configFilePath != "" {
		logger.Printf("Using config file: %s\n", configFilePath)
		config, err = configReader.ConfigFromYaml(configFilePath)
		if err != nil {
			logger.Fatalf("failed to read yaml config: %v", err)
		}
		logger.Printf("Parsed config from file: \n%v", config)

	} else if url != "" && filter != "" {
		config, err = configReader.FilterFromString(filter)
		if err != nil {
			logger.Fatalf("failed to read yaml config: %v", err)
		}
		config.M3uURL = url
		logger.Printf("Parsed config from arguments: \n%v", config)
	} else {
		logger.Fatal("Invalid parameters. Either provide a config file or a URL and filter.")
	}
	return &Organizer{
		logger,
		config,
		m3uParser,
		downloader,
	}
}

func (o *Organizer) Run() {
	m3uFilename := filepath.Join(m3u.DownloadsDir, fmt.Sprintf("%s.m3u", time.Now().Format("2006-01-02")))
	o.downloadM3uFile(m3uFilename)

	m3u, err := o.parser.ParseM3U(m3uFilename)
	if err != nil {
		o.logger.Fatalf("failed to read m3u file: %v", err)
	}

	o.logger.Printf("content of m3u file: %v", m3u)
}

func (o *Organizer) downloadM3uFile(m3uFilename string) {
	err := o.downloader.CreateDownloadsDirIfNotExist()
	if err != nil {
		o.logger.Fatalf("failed to create downloadDir: %v", err)
	}
	err = o.downloader.Download(o.config.M3uURL, m3uFilename)
	if err != nil {
		o.logger.Fatalf("failed to download m3u: %v", err)
	}
}
