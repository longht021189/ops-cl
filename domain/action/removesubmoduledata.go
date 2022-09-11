package action

import (
	"os"
	"path/filepath"

	"github.com/longht021189/ops-cl/model/args"
)

func RemoveSubmoduleData(path string, arg *args.GitArgs) error {
	dir := filepath.Join(*arg.GitRoot, "modules", path)
	return os.RemoveAll(dir)
}
