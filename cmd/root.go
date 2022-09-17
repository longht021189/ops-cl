/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/longht021189/ops-cl/cmd/docker"
	"github.com/longht021189/ops-cl/cmd/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ops-cl",
	Short: "Utitlities",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	git.InitCmd(rootCmd)
	docker.InitCmd(rootCmd)
}
