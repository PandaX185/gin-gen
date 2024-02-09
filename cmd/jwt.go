package cmd

import (
	"github.com/spf13/cobra"
)

var JwtCmd = &cobra.Command{
	Use:   "jwt <path>",
	Short: "generate jwt necessary functions",
	Run: func(cmd *cobra.Command, args []string) {
		CreateJWT(cmd, args)
		CreateJWTMiddleware(cmd, args)
	},
}
