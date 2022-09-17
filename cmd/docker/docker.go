package docker

import (
	"github.com/spf13/cobra"
)

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker Utilities",
}

func InitCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(dockerCmd)
}
