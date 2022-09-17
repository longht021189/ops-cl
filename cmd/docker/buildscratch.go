package docker

import (
	"fmt"
	"os"

	"github.com/longht021189/ops-cl/domain/action"
	"github.com/longht021189/ops-cl/model/args"
	"github.com/spf13/cobra"
)

var buildScratchArgs = &args.BuildScratchArgs{}
var buildScratchCmd = &cobra.Command{
	Use:   "build-scratch",
	Short: "Build Scratch Image from Binary",
	Run: func(cmd *cobra.Command, args []string) {
		err := runBuildScratch(buildScratchArgs)
		if err != nil {
			// TODO Report error
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	dockerCmd.AddCommand(buildScratchCmd)
	buildScratchCmd.Flags().StringSliceVarP(&buildScratchArgs.Binaries, "binary", "b", nil, "Binary Path")
}

func runBuildScratch(data *args.BuildScratchArgs) error {
	for _, binary := range data.Binaries {
		value, err := action.GetBinaryDependencies(binary)

		if err != nil {
			return err
		}

		fmt.Printf(">>>> %v", value)
	}

	return nil
}
