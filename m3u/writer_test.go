package m3u_test

import (
	"os"
	"testing"

	"github.com/fgeck/iptv-m3u-organizer/m3u"
	"github.com/stretchr/testify/assert"
)

func TestM3uWriter_WriteToFile(t *testing.T) {
	writer := m3u.NewWriter()

	tempFile, err := os.CreateTemp("", "m3u_test_*.m3u")
	if err != nil {
		t.Fatalf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	entries := []*m3u.M3UEntry{
		{
			Extinf: &m3u.ExtinfInfo{
				TvgID:      "ZDF.de",
				TvgName:    "DE - ZDF HD",
				TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
				GroupTitle: "|DE| ALLGEMEINES",
				Title:      "DE - ZDF HD",
			},
			URL: "https://somestream.com/stream/abc/def/1234",
		},
		{
			Extinf: &m3u.ExtinfInfo{
				TvgID:      "ARD.de",
				TvgName:    "DE - DAS ERSTE HD",
				TvgLogo:    "https://somelogo.com/daserstehd.png",
				GroupTitle: "|DE| ALLGEMEINES",
				Title:      "DE - DAS ERSTE HD",
			},
			URL: "http://line.dino.ws:80/c23e00bcda/65c1b6b72039/699853",
		},
	}

	err = writer.WriteToFile(entries, tempFile.Name())
	if err != nil {
		t.Fatalf("failed to write to temporary file: %v", err)
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("failed to read temporary file: %v", err)
	}

	expected := `#EXTM3U
#EXTINF:-1 tvg-id="ZDF.de" tvg-name="DE - ZDF HD" tvg-logo="https://somelogo.com/ZDF-HD.jpg" group-title="|DE| ALLGEMEINES",DE - ZDF HD
https://somestream.com/stream/abc/def/1234
#EXTINF:-1 tvg-id="ARD.de" tvg-name="DE - DAS ERSTE HD" tvg-logo="https://somelogo.com/daserstehd.png" group-title="|DE| ALLGEMEINES",DE - DAS ERSTE HD
http://line.dino.ws:80/c23e00bcda/65c1b6b72039/699853
`
	assert.Equal(t, expected, string(data))
}
