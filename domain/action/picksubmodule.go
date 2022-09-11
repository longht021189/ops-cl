package action

import (
	"fmt"

	git "github.com/longht021189/ops-cl/domain/git"
	"github.com/longht021189/ops-cl/model/args"
)

func PickGitSubmodule(data *args.GitArgs) ([]string, error) {
	submodules, err := git.ListSubmodule(data)
	if err != nil {
		return nil, err
	}
	if len(submodules) <= 0 {
		return nil, nil
	}

	fmt.Println("Submodules:")
	for i, submodule := range submodules {
		fmt.Printf("\t%d. %s\n", i+1, submodule)
	}

	var number int
	fmt.Print("Enter number which is index of submodule you'd like to remove: ")
	fmt.Scan(&number)

	return []string{submodules[number-1]}, nil
}
