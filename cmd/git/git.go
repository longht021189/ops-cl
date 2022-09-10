package git

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "git command helpers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("git called")
	},
}

func Get() *cobra.Command {
	return gitCmd
}
