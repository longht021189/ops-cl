package action

import (
	"path/filepath"

	"github.com/longht021189/ops-cl/utils/file"
)

func FindGitRoot(path string) (*string, error) {
	for len(path) > 0 {
		gitRoot := filepath.Join(path, ".git")

		if file.Exists(gitRoot) {
			return &path, nil
		}

		path = filepath.Dir(path)
	}

	return nil, nil
}
