package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fgeck/iptv-m3u-organizer/config"
	"github.com/stretchr/testify/assert"
)

func TestConfigFromString(t *testing.T) {
	configReader := &config.ConfigReader{}

	// Test case: Valid YAML string
	yamlStr := `
m3uURL: https://example.com/playlist.m3u
fullmatch:
  group:
    - Sports
  title:
    - HD Channel
softmatch:
  group:
    - News
`
	cfg, err := configReader.FilterFromString(yamlStr)
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/playlist.m3u", cfg.M3uURL)
	assert.Equal(t, []string{"Sports"}, cfg.FullMatch.Group)
	assert.Equal(t, []string{"HD Channel"}, cfg.FullMatch.Title)
	assert.Equal(t, []string{"News"}, cfg.SoftMatch.Group)

	// Test case: Valid JSON string
	jsonStr := `{"m3uURL":"https://example.com/playlist.m3u","fullmatch":{"group":["Sports"],"title":["HD Channel"]},"softmatch":{"group":["News"]}}`
	cfg, err = configReader.FilterFromString(jsonStr)
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/playlist.m3u", cfg.M3uURL)
	assert.Equal(t, []string{"Sports"}, cfg.FullMatch.Group)
	assert.Equal(t, []string{"HD Channel"}, cfg.FullMatch.Title)
	assert.Equal(t, []string{"News"}, cfg.SoftMatch.Group)

	// Test case: Missing m3uURL field
	invalidStr := `
fullmatch:
  group:
    - Sports
  title:
    - HD Channel
`
	_, err = configReader.FilterFromString(invalidStr)
	assert.EqualError(t, err, "m3uURL is mandatory but not provided")

	// Test case: Invalid string format
	invalidStr = "this is an invalid string"
	_, err = configReader.FilterFromString(invalidStr)
	assert.EqualError(t, err, "failed to parse configuration string as YAML or JSON")
}

func TestConfigFromYaml(t *testing.T) {
	configReader := &config.ConfigReader{}

	// Test case: Valid YAML file
	validFilePath, _ := filepath.Abs("../test/config/valid_config.yaml")
	cfg, err := configReader.ConfigFromYaml(validFilePath)
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com/playlist.m3u", cfg.M3uURL)
	assert.Equal(t, []string{"Sports"}, cfg.FullMatch.Group)
	assert.Equal(t, []string{"HD Channel"}, cfg.FullMatch.Title)
	assert.Equal(t, []string{"News"}, cfg.SoftMatch.Group)

	// Test case: File not found
	_, err = configReader.ConfigFromYaml("nonexistent.yaml")
	assert.Error(t, err)
	assert.True(t, os.IsNotExist(err))

	// Test case: Invalid YAML file
	invalidFilePath, _ := filepath.Abs("testdata/invalid_config.yaml")
	_, err = configReader.ConfigFromYaml(invalidFilePath)
	assert.Error(t, err)
}
