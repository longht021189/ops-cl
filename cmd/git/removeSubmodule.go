package git

import (
	"fmt"
	"log"
	"os"

	domain "github.com/longht021189/ops-cl/domain/action"
	"github.com/spf13/cobra"
)

var (
	submodulePaths []string
)

var removeSubmoduleCmd = &cobra.Command{
	Use:   "remove-submodule",
	Short: "Remove Git Submodule",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// TODO input dir first
		targetDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		if len(submodulePaths) <= 0 {
			submodulePaths, err = domain.PickGitSubmodule()
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\n\naaaa: %v", submodulePaths)
	},
}

func init() {
	gitCmd.AddCommand(removeSubmoduleCmd)
	removeSubmoduleCmd.Flags().StringSliceVar(&submodulePaths, "path", nil, "Submodule Path")
}

// git rm -f projects/github.com/longht021189/ops-cl
// rm -rf .git/modules/projects/github.com/longht021189/ops-cl
// git config --remove-section submodule.projects/github.com/longht021189/ops-cl
