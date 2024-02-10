package cmd

import (
	"github.com/spf13/cobra"
)

var RepoCmd = &cobra.Command{
	Use:   "repo <path>",
	Short: "Create repository file",
	Run: func(cmd *cobra.Command, args []string) {
		CreateModel(cmd, args)
		CreateRepo(cmd, args)
	},
}
