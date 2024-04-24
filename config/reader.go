package config

import (
	"encoding/json"
	"errors"
	"os"

	"gopkg.in/yaml.v2"
)

type ConfigReader struct{}

func NewConfigReader() *ConfigReader {
	return &ConfigReader{}
}

func (y *ConfigReader) ConfigFromYaml(filePath string) (*Config, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.UnmarshalStrict(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (y *ConfigReader) FilterFromString(filterString string) (*Config, error) {
	var cfg Config

	err := yaml.Unmarshal([]byte(filterString), &cfg)
	if err == nil {
		if cfg.M3uURL == "" {
			return nil, errors.New("m3uURL is mandatory but not provided")
		}
		return &cfg, nil
	}

	err = json.Unmarshal([]byte(filterString), &cfg)
	if err == nil {
		if cfg.M3uURL == "" {
			return nil, errors.New("m3uURL is mandatory but not provided")
		}
		return &cfg, nil
	}

	return nil, errors.New("failed to parse configuration string as YAML or JSON")
}
