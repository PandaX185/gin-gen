package cmd

import "github.com/spf13/cobra"

var ControllerCmd = &cobra.Command{
	Use:   "controller <path>",
	Short: "Create controller file",
	Run: func(cmd *cobra.Command, args []string) {
		CreateModel(cmd, args)
		CreateRepo(cmd, args)
		CreateService(cmd, args)
		CreateController(cmd, args)
	},
}
