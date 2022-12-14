package git

import (
	a "github.com/longht021189/ops-cl/model/args"
	"github.com/longht021189/ops-cl/utils/exe"
)

func RemoveSubmodule(path string, data *a.GitArgs) error {
	args := data.BuildPrefixCmd()
	args = append(args, "rm", "-f", path)

	_, err := exe.Run(args...)
	if err != nil {
		return err
	}

	return nil
}
