package cmd

import (
	"github.com/spf13/cobra"
)

var ModelCmd = &cobra.Command{
	Use:   "model <path>",
	Short: "Create model file",
	Run: func(cmd *cobra.Command, args []string) {
		CreateModel(cmd, args)
	},
}
