/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// adminCmd represents the admin command
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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

		fmt.Printf("admin called %v, %v", ex, string(out))

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adminCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// adminCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}