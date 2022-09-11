package git

import (
	"strings"

	"github.com/longht021189/ops-cl/model/args"
	"github.com/longht021189/ops-cl/utils/exe"
)

func ListSubmodule(data *args.GitArgs) ([]string, error) {
	args := data.BuildPrefixCmd()
	args = append(args, "submodule")

	output, err := exe.Run(args...)
	if err != nil {
		return nil, err
	}

	var result []string

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		cols := strings.Split(line, " ")

		var columns []string
		for _, c := range cols {
			data := strings.TrimSpace(c)
			if len(data) > 0 {
				columns = append(columns, data)
			}
		}

		if len(columns) < 2 {
			continue
		}

		result = append(result, columns[1])
	}

	return result, nil
}
