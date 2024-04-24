package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fgeck/iptv-m3u-organizer/download"
	"github.com/fgeck/iptv-m3u-organizer/m3u"
	"github.com/fgeck/iptv-m3u-organizer/yaml"
)

func main() {
	logger := log.New(os.Stdout, "main: ", log.Ldate|log.Ltime|log.Lshortfile)
	yamlReader := &yaml.YamlReader{}
	m3uParser := &m3u.M3uParser{}
	downloader := download.NewDownloader()
	m3uFilename := filepath.Join(download.DownloadsDir, fmt.Sprintf("%s.m3u", time.Now().Format("2006-01-02")))
	config, err := yamlReader.ReadConfig("config.yaml")
	if err != nil {
		logger.Fatalf("failed to read yaml config: %v", err)
	}
	logger.Printf("content of yaml config: %v", config)
	err = downloader.CreateDownloadsDirIfNotExist()
	if err != nil {
		logger.Fatalf("failed to create downloadDir: %v", err)
	}
	err = downloader.Download(config.M3uURL, m3uFilename)
	if err != nil {
		logger.Fatalf("failed to download m3u: %v", err)
	}
	m3u, err := m3uParser.ParseM3U(m3uFilename)
	if err != nil {
		logger.Fatalf("failed to read m3u file: %v", err)
	}
	logger.Printf("content of m3u file: %v", m3u)
}
