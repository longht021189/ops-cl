package action

import (
	"os"
	"path/filepath"

	"github.com/longht021189/ops-cl/model/args"
	"github.com/longht021189/ops-cl/utils/file"
)

func RemoveSubmoduleData(path string, arg *args.GitArgs) error {
	dir := filepath.Join(*arg.GitRoot, "modules", path)

	if file.Exists(dir) {
		return os.RemoveAll(dir)
	}

	return nil
}
