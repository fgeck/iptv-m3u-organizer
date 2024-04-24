package yaml

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	M3uURL    string        `yaml:"m3uURL"`
	FullMatch MatchCriteria `yaml:"fullmatch"`
	SoftMatch MatchCriteria `yaml:"softmatch"`
}

type MatchCriteria struct {
	Group []string `yaml:"group"`
	Title []string `yaml:"title"`
}

type YamlReader struct{}

func (y *YamlReader) ReadConfig(filename string) (*Config, error) {
	// Read YAML file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal YAML into Config struct
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
