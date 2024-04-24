package m3u

import (
	"testing"

	"github.com/fgeck/iptv-m3u-organizer/config"
	"github.com/stretchr/testify/assert"
)

func TestM3uFilter_Filter_fullSoftGroupNameMatching(t *testing.T) {
	config := &config.Config{
		M3uURL: "",
		FullMatch: config.MatchCriteria{
			Group: []string{"|DE| ALLGEMEINES"},
			Name:  []string{"DE - ZDF HD"},
		},
		SoftMatch: config.MatchCriteria{
			Group: []string{"|DE| Finance"},
			Name:  []string{"DW News"},
		},
	}

	m3uEntries := []*M3UEntry{
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
		{
			Extinf: &ExtinfInfo{
				TvgID:      "ZDF.de",
				TvgName:    "DE - ZDF HD",
				TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
				GroupTitle: "|DE| Not-So-General",
				Title:      "DE - ZDF HD",
			},
			URL: "https://somestream.com/stream/abc/def/1434",
		},
		{
			Extinf: &ExtinfInfo{
				TvgID:      "DW.de",
				TvgName:    "DW News",
				TvgLogo:    "https://somelogo.com/dw-news.png",
				GroupTitle: "|DE| ALLGEMEINES",
				Title:      "DW News",
			},
			URL: "https://somestream.com/stream/abc/def/9874",
		},
		{
			Extinf: &ExtinfInfo{
				TvgID:      "bloomberg.de",
				TvgName:    "bloomberg News",
				TvgLogo:    "https://somelogo.com/bloomberg-news.png",
				GroupTitle: "|DE| Finance",
				Title:      "bloomberg News",
			},
			URL: "https://somestream.com/stream/abc/def/9837",
		},
		{
			Extinf: &ExtinfInfo{
				TvgID:      "MTV.de",
				TvgName:    "Mtv Music",
				TvgLogo:    "https://somelogo.com/mtv-music.png",
				GroupTitle: "|DE| Music",
				Title:      "Mtv Music",
			},
			URL: "https://somestream.com/stream/abc/def/9879",
		},
	}

	expected := []*M3UEntry{
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
		{
			Extinf: &ExtinfInfo{
				TvgID:      "ZDF.de",
				TvgName:    "DE - ZDF HD",
				TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
				GroupTitle: "|DE| Not-So-General",
				Title:      "DE - ZDF HD",
			},
			URL: "https://somestream.com/stream/abc/def/1434",
		},
		{
			Extinf: &ExtinfInfo{
				TvgID:      "DW.de",
				TvgName:    "DW News",
				TvgLogo:    "https://somelogo.com/dw-news.png",
				GroupTitle: "|DE| ALLGEMEINES",
				Title:      "DW News",
			},
			URL: "https://somestream.com/stream/abc/def/9874",
		},
		{
			Extinf: &ExtinfInfo{
				TvgID:      "bloomberg.de",
				TvgName:    "bloomberg News",
				TvgLogo:    "https://somelogo.com/bloomberg-news.png",
				GroupTitle: "|DE| Finance",
				Title:      "bloomberg News",
			},
			URL: "https://somestream.com/stream/abc/def/9837",
		},
	}

	filter := M3uFilter{}
	filteredEntries, err := filter.Filter(config, m3uEntries)
	assert.NoError(t, err)

	compareEntries(t, filteredEntries, expected)
}

func TestM3uFilter_Filter_with_duplicates(t *testing.T) {

	config := &config.Config{
		M3uURL: "",
		FullMatch: config.MatchCriteria{
			Group: []string{"|DE| ALLGEMEINES"},
			Name:  []string{"DE - DAS ERSTE HD"},
		},
		SoftMatch: config.MatchCriteria{
			Group: []string{},
			Name:  []string{"ZDF HD"},
		},
	}
	m3uEntries := []*M3UEntry{
		{
			Extinf: &ExtinfInfo{
				TvgID:      "ZDF.de",
				TvgName:    "DE - ZDF HD",
				TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
				GroupTitle: "|DE| Not-So-General",
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
		{
			Extinf: &ExtinfInfo{
				TvgID:      "ZDF.de",
				TvgName:    "DE - ZDF HD",
				TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
				GroupTitle: "|DE| Not-So-General",
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
	}

	expected := []*M3UEntry{
		{
			Extinf: &ExtinfInfo{
				TvgID:      "ZDF.de",
				TvgName:    "DE - ZDF HD",
				TvgLogo:    "https://somelogo.com/ZDF-HD.jpg",
				GroupTitle: "|DE| Not-So-General",
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
	}

	filter := M3uFilter{}
	filteredEntries, err := filter.Filter(config, m3uEntries)
	assert.NoError(t, err)

	compareEntries(t, filteredEntries, expected)
}

func compareEntries(t *testing.T, got, want []*M3UEntry) {
	if len(got) != len(want) {
		t.Errorf("Length of filtered entries does not match expected: got %d, want %d", len(got), len(want))
		return
	}

	// Compare each entry
	for i, ge := range got {
		we := want[i]
		if ge.Extinf.TvgName != we.Extinf.TvgName || ge.Extinf.TvgLogo != we.Extinf.TvgLogo {
			t.Errorf("Mismatched Extinf: got %v, want %v", ge.Extinf, we.Extinf)
		}
		if ge.URL != we.URL {
			t.Errorf("Mismatched URL: got %v, want %v", ge.URL, we.URL)
		}
	}
}
