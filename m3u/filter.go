package m3u

import (
	"strings"

	"github.com/fgeck/iptv-m3u-organizer/config"
)

type M3uFilter struct{}

func (f *M3uFilter) Filter(config *config.Config, m3uEntries []*M3UEntry) ([]*M3UEntry, error) {
	filteredEntries := make([]*M3UEntry, 0, len(m3uEntries))
	deduplicatedEntries := make(map[string]int)

	for _, entry := range m3uEntries {
		hash, err := entry.Hash()
		if err != nil {
			return nil, err
		}
		_, exists := deduplicatedEntries[hash]
		shouldAdd := f.shouldAdd(entry, config)
		if shouldAdd && !exists {
			deduplicatedEntries[hash] = 1
			filteredEntries = append(filteredEntries, entry)
		}

	}

	return filteredEntries, nil
}

func (f *M3uFilter) shouldAdd(entry *M3UEntry, config *config.Config) bool {
	for _, criteria := range config.FullMatch.Group {
		if entry.Extinf.GroupTitle == criteria {
			return true
		}
	}
	for _, criteria := range config.FullMatch.Name {
		if entry.Extinf.TvgName == criteria {
			return true
		}
	}

	for _, criteria := range config.SoftMatch.Group {
		if strings.Contains(entry.Extinf.GroupTitle, criteria) {
			return true
		}

	}
	for _, criteria := range config.SoftMatch.Name {
		if strings.Contains(entry.Extinf.TvgName, criteria) {
			return true
		}

	}
	return false
}
