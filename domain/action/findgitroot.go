package action

import (
	"os"
	"path/filepath"
)

func FindGitRoot(path string) (*string, error) {
	for len(path) > 0 {
		gitRoot := filepath.Join(path, ".git")

		if _, err := os.Stat(gitRoot); !os.IsNotExist(err) {
			return &path, nil
		}

		path = filepath.Dir(path)
	}

	return nil, nil
}
