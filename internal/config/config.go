package config

import (
	"errors"
	"gopkg.in/yaml.v2"
)

// Config represents the configuration for the application.
type Config struct {
	// SrcGlobPatterns are globs to apply when searching for source files.
	SrcGlobPatterns []string `yaml:"src_glob_patterns"`
	// WarningAgeDays is the number of days after which an unresolved TODO will result in a warning.
	WarningAgeDays int `yaml:"warning_age_days"`
	// JIRAAddress is the address to use for JIRA API calls.
	JIRAAddress string `yaml:"jira_address"`
	// GithubRepoAddress is the address used to make links to LOCs.
	GithubRepoAddress string `yaml:"github_repo_address"`
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
	if c.WarningAgeDays <= 0 {
		return errors.New("config is invalid: no warning_age_days given or non-positive value given")
	}
	if c.JIRAAddress == "" {
		return errors.New("config is invalid: no jira_address given")
	}
	if c.GithubRepoAddress == "" {
		return errors.New("config is invalid: no github_repo_address given")
	}
	return nil
}
