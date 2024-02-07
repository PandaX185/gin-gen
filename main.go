package main

import (
	"github/PandaX185/gin-gen/cmd"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gin-gen",
	Version: "1.0.0",
	Short:   "Gin Gen is used to automate backend development",
}

func main() {
	rootCmd.AddCommand(cmd.JwtCmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
