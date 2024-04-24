package m3u_test

import (
	"testing"

	"github.com/fgeck/iptv-m3u-organizer/m3u"
	"github.com/stretchr/testify/assert"
)

func TestM3UEntry_String(t *testing.T) {
	entry := &m3u.M3UEntry{
		Extinf: &m3u.ExtinfInfo{
			TvgID:      "ZDF.de",
			TvgName:    "DE - ZDF HD",
			TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
			GroupTitle: "|DE| ALLGEMEINES",
			Title:      "DE - ZDF HD",
		},
		URL: "https://somestream.com/stream/abc/def/1234",
	}

	expected := `#EXTINF:-1 tvg-id="ZDF.de" tvg-name="DE - ZDF HD" tvg-logo="https://somelogo.com/ZDF-HD.jpg" group-title="|DE| ALLGEMEINES",DE - ZDF HD
https://somestream.com/stream/abc/def/1234`

	assert.Equal(t, expected, entry.String())
}
