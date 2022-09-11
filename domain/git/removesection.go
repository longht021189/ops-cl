package git

import (
	"fmt"

	a "github.com/longht021189/ops-cl/model/args"
	"github.com/longht021189/ops-cl/utils/exe"
)

func RemoveSection(path string, data *a.GitArgs) error {
	args := data.BuildPrefixCmd()
	args = append(args, "config", "--remove-section", fmt.Sprintf("submodule.%s", path))

	_, err := exe.Run(args...)
	if err != nil {
		return err
	}

	return nil
}
