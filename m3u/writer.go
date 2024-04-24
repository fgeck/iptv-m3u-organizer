package m3u

import (
	"log"
	"os"
)

const (
	startingLine = "#EXTM3U"
)

type M3uWriter struct {
	logger *log.Logger
}

func NewWriter() *M3uWriter {
	return &M3uWriter{
		log.New(os.Stdout, "organizer: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Takes a list of M3UEntry and writes it to a file
// It is expected to be written into following format:
// --------------------------------------------------------------------------------------------------------------------------------------------
// #EXTM3U
// #EXTINF:-1 tvg-id="ZDF.de" tvg-name="DE - ZDF HD" tvg-logo="https://somelogo.com/ZDF-HD.jpg" group-title="|DE| ALLGEMEINES",DE - ZDF HD
// https://somestream.com/stream/abc/def/1234
// --------------------------------------------------------------------------------------------------------------------------------------------
func (w *M3uWriter) WriteToFile(content []*M3UEntry, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(startingLine + "\n")
	if err != nil {
		return err
	}

	for _, entry := range content {
		_, err = file.WriteString(entry.String() + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
