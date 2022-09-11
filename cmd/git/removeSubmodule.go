package git

import (
	"errors"
	"log"
	"os"

	domain "github.com/longht021189/ops-cl/domain/action"
	"github.com/longht021189/ops-cl/domain/git"
	a "github.com/longht021189/ops-cl/model/args"
	"github.com/spf13/cobra"
)

var (
	gitArgs = &a.GitArgs{}
)

var removeSubmoduleCmd = &cobra.Command{
	Use:   "remove-submodule",
	Short: "Remove Git Submodule",
	Run: func(cmd *cobra.Command, args []string) {
		err := runRemoveSubmodule()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	gitCmd.AddCommand(removeSubmoduleCmd)
	removeSubmoduleCmd.Flags().StringSliceVar(&gitArgs.SubmodulePaths, "path", nil, "Submodule Path")
}

func runRemoveSubmodule() error {
	var err error

	// TODO input dir first
	gitArgs.TargetDir, err = os.Getwd()
	if err != nil {
		return err
	}

	gitArgs.GitRoot, err = domain.FindGitRoot(gitArgs.TargetDir)
	if err != nil {
		return err
	}
	if gitArgs.GitRoot == nil {
		// TODO define error code
		return errors.New("Git Root not Found")
	}

	if len(gitArgs.SubmodulePaths) <= 0 {
		gitArgs.SubmodulePaths, err = domain.PickGitSubmodule(gitArgs)
	}
	if err != nil {
		log.Fatal(err)
	}

	for _, submodule := range gitArgs.SubmodulePaths {
		err = git.RemoveSubmodule(submodule, gitArgs)
		if err != nil {
			return err
		}

		err = git.RemoveSection(submodule, gitArgs)
		if err != nil {
			return err
		}

		err = domain.RemoveSubmoduleData(submodule, gitArgs)
		if err != nil {
			return err
		}
	}

	return nil
}
