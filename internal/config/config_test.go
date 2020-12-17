package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWellFormedConfigIsLoadedSuccessfully(t *testing.T) {
	data := `
src_glob_paths: ["path/*", "another/path/*.ext"]
github_url: "github.com/neprune/todo"
`
	c, err := LoadFromYAMLData([]byte(data))
	require.NoError(t, err)
	require.Equal(t, "github.com/neprune/todo", c.GithubURL)
	require.Contains(t, c.SrcGlobPaths, "path/*")
	require.Contains(t, c.SrcGlobPaths, "another/path/*.ext")
	require.Equal(t, 2, len(c.SrcGlobPaths))
}

func TestConfigWithoutSrcGlobPathsFailsToLoad(t *testing.T) {
	data := `
github_url: "github.com/neprune/todo"
`
	_, err := LoadFromYAMLData([]byte(data))
	require.Error(t, err)
}