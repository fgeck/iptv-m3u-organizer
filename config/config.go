package config

type Config struct {
	M3uURL    string        `yaml:"m3uURL"`
	FullMatch MatchCriteria `yaml:"fullmatch"`
	SoftMatch MatchCriteria `yaml:"softmatch"`
}

type MatchCriteria struct {
	Group []string `yaml:"group"` // group-title / category
	Name []string `yaml:"name"` // tvg-name
}
