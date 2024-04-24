package m3u

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// ExtinfInfo represents parsed information from #EXTINF line
type ExtinfInfo struct {
	TvgID       string
	TvgName     string
	TvgLogo     string
	GroupTitle  string
	ChannelName string
}

// M3UEntry represents an entry in the M3U playlist containing ExtinfInfo and URL
type M3UEntry struct {
	Extinf ExtinfInfo
	URL    string
}

type M3uParser struct{}

// parseM3U parses the M3U playlist file and returns a slice of M3UEntry
func (m *M3uParser) ParseM3U(filename string) ([]M3UEntry, error) {
	var entries []M3UEntry

	// Open the .m3u file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	var extinf *ExtinfInfo
	// Parse each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\\", "")
		if strings.HasPrefix(line, "#EXTINF:") {
			extinf = m.parseExtinf(line)
		} else if extinf != nil {
			entries = append(entries, M3UEntry{Extinf: *extinf, URL: line})
			extinf = nil
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

// parseExtinf parses a line of #EXTINF and returns ExtinfInfo
func (m *M3uParser) parseExtinf(line string) *ExtinfInfo {
	// Check if the line starts with #EXTINF
	if !strings.HasPrefix(line, "#EXTINF:") {
		return nil
	}

	// Split the line by comma and space
	parts := strings.SplitN(line[8:], ",", 2)
	if len(parts) != 2 {
		return nil
	}

	// Extract attributes
	attributeRegex := regexp.MustCompile(`(\w+)="([^"]*)"`)
	attributes := make(map[string]string)
	matches := attributeRegex.FindAllStringSubmatch(parts[1], -1)
	for _, match := range matches {
		attributes[match[1]] = match[2]
	}
//	"#EXTINF:-1 tvg-id=\"\" tvg-name=\"##### GENERAL #####\" tvg-logo=\"http://logo.protv.cc/picons/logos/france/FRANCE.png\" group-title=\"|EU| FRANCE HEVC\",##### GENERAL #####"
	return &ExtinfInfo{
		TvgID:       attributes["tvg-id"],
		TvgName:     attributes["tvg-name"],
		TvgLogo:     attributes["tvg-logo"],
		GroupTitle:  attributes["group-title"],
		ChannelName: strings.TrimSpace(parts[1]),
	}
}
