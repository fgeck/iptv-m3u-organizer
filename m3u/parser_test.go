package m3u

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseExtinfLine(t *testing.T) {
	testCases := []struct {
		name     string
		line     string
		expected *ExtinfInfo
		err      bool
	}{
		{
			name: "Valid #EXTINF line",
			line: "#EXTINF:-1 tvg-id=\"ZDF.de\" tvg-name=\"DE - ZDF HD\" tvg-logo=\"https://somelogo.com/ZDF-HD.jpg\" group-title=\"|DE| ALLGEMEINES\",DE - ZDF HD",
			expected: &ExtinfInfo{
				TvgID:      "ZDF.de",
				TvgName:    "DE - ZDF HD",
				TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
				GroupTitle: "|DE| ALLGEMEINES",
				Title:      "DE - ZDF HD",
			},
			err: false,
		},
		{
			name:     "Invalid #EXTINF line",
			line:     "#EXTINF:",
			expected: nil,
			err:      false,
		},
	}

	parser := NewM3uParser()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info := parser.parseExtinfLine(tc.line)
			assert.Equal(t, tc.expected, info)
		})
	}
}

func TestParseM3U(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		expected []*M3UEntry
		err      bool
	}{
		{
			name:     "Valid M3U file",
			filename: "../test/m3u/valid.m3u",
			expected: []*M3UEntry{
				{
					Extinf: &ExtinfInfo{
						TvgID:      "",
						TvgName:    "##### DE ALLGEMEINES #####",
						TvgLogo:    "https://somelogo.com/germany/germany_640.png",
						GroupTitle: "|DE| ALLGEMEINES",
						Title:      "##### DE ALLGEMEINES #####",
					},
					URL: "https://somestream.com/stream/abc/def/123",
				},
				{
					Extinf: &ExtinfInfo{
						TvgID:      "ZDF.de",
						TvgName:    "DE - ZDF HD",
						TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
						GroupTitle: "|DE| ALLGEMEINES",
						Title:      "DE - ZDF HD",
					},
					URL: "https://somestream.com/stream/abc/def/1234",
				},
				{
					Extinf: &ExtinfInfo{
						TvgID:      "ARD.de",
						TvgName:    "DE - DAS ERSTE HD",
						TvgLogo:    "https://somelogo.com/daserstehd.png",
						GroupTitle: "|DE| ALLGEMEINES",
						Title:      "DE - DAS ERSTE HD",
					},
					URL: "https://somestream.com/stream/abc/def/12345",
				},
			},
			err: false,
		},
		{
			name:     "File not found",
			filename: "testdata/nonexistent.m3u",
			expected: nil,
			err:      true,
		},
	}
	parser := NewM3uParser()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			entries, err := parser.ParseM3U(tc.filename)

			if tc.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, entries)
			}
		})
	}
}
