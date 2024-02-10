package cmd

import (
	"github.com/spf13/cobra"
)

var ServiceCmd = &cobra.Command{
	Use:   "service <path>",
	Short: "Create service",
	Run: func(cmd *cobra.Command, args []string) {
		CreateModel(cmd, args)
		CreateRepo(cmd, args)
		CreateService(cmd, args)
	},
}
