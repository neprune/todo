package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
)

func GetCommit(dir string) (string, error) {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return "", fmt.Errorf("failed to open repo dir: %w", err)
	}
	h, err := repo.Head()
	if err != nil {
		return "", fmt.Errorf("failed to get head: %w", err)
	}
	return h.Hash().String(), nil
}
