package github

import "fmt"

func GenerateLOCGithubURL(repoURL string, commit string, filepath string, lineNumber int) string {
	return fmt.Sprintf("%sblob/%s/%s#L%d", repoURL, commit, filepath, lineNumber)
}
