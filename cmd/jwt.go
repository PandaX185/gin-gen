package cmd

import (
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var JwtCmd = &cobra.Command{
	Use:   "jwt <path>",
	Short: "generate login with jwt token",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			file *os.File
			err  error
		)
		if len(args) == 0 {
			file, err = CreateJwtFile()
		} else {
			file, err = CreateJwtFile(args[0])
		}
		if err != nil {
			panic(err)
		}
		jwtFile, _ := os.Open("templates/jwt.txt")
		jwtTemplate, _ := io.ReadAll(jwtFile)
		file.Write(jwtTemplate)
		exec.Command("go", "mod", "tidy").Run()
		defer func() {
			file.Close()
			jwtFile.Close()
		}()
	},
}
