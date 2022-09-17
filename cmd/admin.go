package cmd

import (
	"github.com/spf13/cobra"
)

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		/* Linux/OSX,...
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}

		argsStr := strings.Join(args, " ")
		cmaad := exec.Command("/bin/sh", "-c", fmt.Sprintf("sudo %s %s", ex, argsStr))
		out, err := cmaad.Output()

		if err != nil {
			panic(err)
		}

		fmt.Printf("admin called %v, %v", ex, string(out))*/

		// cmaad := exec.Command("/bin/sh", "-c", "sudo ls")
		// out, err := cmaad.Output()
		// fmt.Printf("admin called %v, %v", out, err)

		// cmaad = exec.Command("/bin/sh", "-c", "sudo ls")
		// out, err = cmaad.Output()
		// fmt.Printf("admin called %v, %v", out, err)
	},
}

func init() {
	rootCmd.AddCommand(adminCmd)
}
