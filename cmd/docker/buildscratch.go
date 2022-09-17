package docker

import (
	"fmt"
	"os"
	t "text/template"

	_ "embed"

	"github.com/longht021189/ops-cl/domain/action"
	"github.com/longht021189/ops-cl/model/args"
	"github.com/longht021189/ops-cl/utils/file"
	"github.com/longht021189/ops-cl/utils/slice"
	"github.com/spf13/cobra"
)

//go:embed scratch.dockerfile
var scratchTemplate string
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
	buildScratchCmd.Flags().StringVarP(&buildScratchArgs.Dockerfile, "output", "o", "Dockerfile", "Output")
	buildScratchCmd.Flags().StringVar(&buildScratchArgs.BuilderName, "builder", "builder", "Docker Builder Name")
}

func runBuildScratch(data *args.BuildScratchArgs) error {
	var needChecks []string
	for _, f := range data.Binaries {
		if !file.Exists(f) {
			return fmt.Errorf("asdasdasdsads") // TODO FIXME
		}

		needChecks = append(needChecks, f)
	}

	var files []string
	var index = 0
	for index < len(needChecks) {
		path := needChecks[index]
		values, err := action.GetBinaryDependencies(path)

		if err != nil {
			return err
		}

		files = append(files, path)

		for _, value := range values {
			if !slice.Contains(value, files) {
				needChecks = append(needChecks, value)
			}
		}

		index += 1
	}

	template, err := t.New("scratch-docker").Parse(scratchTemplate)
	if err != nil {
		return err
	}

	file, err := os.Create(data.Dockerfile)
	if err != nil {
		return err
	}

	defer file.Close()

	d := struct {
		Files   []string
		Builder string
	}{
		Files:   files,
		Builder: data.BuilderName,
	}

	err = template.Execute(file, d)
	if err != nil {
		return err
	}

	return nil
}
