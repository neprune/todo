package config

import (
	"errors"
	"gopkg.in/yaml.v2"
)

// Config represents the configuration for the application.
type Config struct {
	// SrcGlobPatterns are globs to apply when searching for source files.
	SrcGlobPatterns []string `yaml:"src_glob_patterns"`
	// GithubURL is the URL of the github repo, used for generating links to the TODO lines.
	// Optional.
	GithubURL string `yaml:"github_url"`
}

// LoadFromYAMLData loads a config from the given YAML data, returning an error if it fails or is invalid.
func LoadFromYAMLData(y []byte) (*Config, error) {
	c := Config{}
	err := yaml.Unmarshal([]byte(y), &c)
	if err != nil {
		return nil, err
	}
	return &c, c.valid()
}

func (c *Config) valid() error {
	if c.SrcGlobPatterns == nil {
		return errors.New("config is invalid: no src_glob_paths given")
	}
	return nil
}
