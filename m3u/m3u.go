package m3u

import (
	"encoding/json"
	"fmt"
)

// ExtinfInfo represents parsed information from #EXTINF line
type ExtinfInfo struct {
	TvgID      string // tvg-id
	TvgName    string // tvg-name
	TvgLogo    string // tvg-logo
	GroupTitle string // group-title
	Title      string // title
}

// M3UEntry represents an entry in the M3U playlist containing ExtinfInfo and URL
type M3UEntry struct {
	Extinf *ExtinfInfo
	URL    string
}

// Returns a 2 line string representation of the M3UEntry in following format:
// --------------------------------------------------------------------------------------------------------------------------------------------
// #EXTINF:-1 tvg-id="ZDF.de" tvg-name="DE - ZDF HD" tvg-logo="https://somelogo.com/ZDF-HD.jpg" group-title="|DE| ALLGEMEINES",DE - ZDF HD
// https://somestream.com/stream/abc/def/1234
// --------------------------------------------------------------------------------------------------------------------------------------------
func (e *M3UEntry) String() string {
	return fmt.Sprintf("#EXTINF:-1 tvg-id=\"%s\" tvg-name=\"%s\" tvg-logo=\"%s\" group-title=\"%s\",%s", e.Extinf.TvgID, e.Extinf.TvgName, e.Extinf.TvgLogo, e.Extinf.GroupTitle, e.Extinf.Title) + "\n" + e.URL
}

func (e *M3UEntry) Hash() (string, error) {
	bytes, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	hash := fmt.Sprintf("%x", bytes)
	return hash, nil
}
