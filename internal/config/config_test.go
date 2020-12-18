package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWellFormedConfigIsLoadedSuccessfully(t *testing.T) {
	data := `
src_glob_patterns: ["path/*", "another/path/*.ext"]
warning_age_days: 5
jira_address: "https://my.jira.com"
github_repo_address: "https://github.com/neprune/todo/"
`
	c, err := LoadFromYAMLData([]byte(data))
	require.NoError(t, err)
	require.Contains(t, c.SrcGlobPatterns, "path/*")
	require.Contains(t, c.SrcGlobPatterns, "another/path/*.ext")
	require.Equal(t, 2, len(c.SrcGlobPatterns))
	require.Equal(t, 5, c.WarningAgeDays)
	require.Equal(t, "https://my.jira.com", c.JIRAAddress)
	require.Equal(t, "https://github.com/neprune/todo/", c.GithubRepoAddress)
}

func TestConfigWithoutSrcGlobPathsFailsToLoad(t *testing.T) {
	data := `
warning_age_days: 5
jira_address: "https://my.jira.com"
github_repo_address: "https://github.com/neprune/todo/"
`
	_, err := LoadFromYAMLData([]byte(data))
	require.Error(t, err)
}

func TestConfigWithoutWarningAgeDaysFailsToLoad(t *testing.T) {
	data := `
src_glob_patterns: ["path/*", "another/path/*.ext"]
jira_address: "https://my.jira.com"
github_repo_address: "https://github.com/neprune/todo/"
`
	_, err := LoadFromYAMLData([]byte(data))
	require.Error(t, err)
}
func TestConfigWithoutJIRAAddressFailsToLoad(t *testing.T) {
	data := `
src_glob_patterns: ["path/*", "another/path/*.ext"]
warning_age_days: 5
github_repo_address: "https://github.com/neprune/todo/"
`
	_, err := LoadFromYAMLData([]byte(data))
	require.Error(t, err)
}

func TestConfigWithoutGithubRepoAddressFailsToLoad(t *testing.T) {
	data := `
src_glob_patterns: ["path/*", "another/path/*.ext"]
warning_age_days: 5
jira_address: "https://my.jira.com"
`
	_, err := LoadFromYAMLData([]byte(data))
	require.Error(t, err)
}
