package docker

import (
	"fmt"

	"github.com/longht021189/ops-cl/model/args"
	"github.com/spf13/cobra"
)

var buildScratchArgs = &args.BuildScratchArgs{}
var buildScratchCmd = &cobra.Command{
	Use:   "build-scratch",
	Short: "Build Scratch Image from Binary",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("docker2 called")
	},
}

func init() {
	dockerCmd.AddCommand(buildScratchCmd)
	buildScratchCmd.Flags().StringSliceVarP(&buildScratchArgs.Binaries, "binary", "b", nil, "Binary Path")
}
