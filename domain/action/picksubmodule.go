package action

import (
	"fmt"

	git "github.com/longht021189/ops-cl/domain/git"
)

func PickGitSubmodule() ([]string, error) {
	submodules, err := git.ListSubmodule()
	if err != nil {
		return nil, err
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
