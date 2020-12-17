package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWellFormedConfigIsLoadedSuccessfully(t *testing.T) {
	data := `
src_glob_patterns: ["path/*", "another/path/*.ext"]
`
	c, err := LoadFromYAMLData([]byte(data))
	require.NoError(t, err)
	require.Contains(t, c.SrcGlobPatterns, "path/*")
	require.Contains(t, c.SrcGlobPatterns, "another/path/*.ext")
	require.Equal(t, 2, len(c.SrcGlobPatterns))
}

func TestConfigWithoutSrcGlobPathsFailsToLoad(t *testing.T) {
	data := ``
	_, err := LoadFromYAMLData([]byte(data))
	require.Error(t, err)
}
