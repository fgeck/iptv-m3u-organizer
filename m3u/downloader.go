package m3u

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/fgeck/iptv-m3u-organizer/config"
)

const (
	DownloadsDir = "./downloads"
)

type Downloader struct {
	logger *log.Logger
}

func NewDownloader() *Downloader {
	return &Downloader{
		logger: log.New(os.Stdout, "downloader: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (d *Downloader) Download(m3uUrl string, filename string, auth *config.AuthenticationInformation) error {
	// Check if the file exists
	if _, err := os.Stat(filename); err == nil {
		d.logger.Println("File already exists. Skip download")
		return nil
	}

	parsedURL, err := url.Parse(m3uUrl)
	if err != nil {
		return err
	}

	switch auth.AuthType {
	case config.BasicAuth:
		d.logger.Fatal("Basic authentication is not supported yet")
	case config.URLParamAuth:
		// params := parsedURL.
		parsedURL.RawQuery += fmt.Sprintf("&username=%s&password=%s", auth.User, auth.Password)
	}

	err = d.downloadToFile(parsedURL.String(), filename)
	if err != nil {
		d.logger.Printf("failed to create download file: %v", err)
		return err
	}

	return nil
}

func (d *Downloader) downloadToFile(url string, filePath string) error {
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (d *Downloader) CreateDownloadsDirIfNotExist() error {
	if _, err := os.Stat(DownloadsDir); os.IsNotExist(err) {
		return os.Mkdir(DownloadsDir, 0755)
	}
	return nil
}
