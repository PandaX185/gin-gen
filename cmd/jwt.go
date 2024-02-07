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
			file, err = CreateFile()
		} else {
			file, err = CreateFile(args[0])
		}
		if err != nil {
			panic(err)
		}
		defer func() {
			err := file.Close()
			if err != nil {
				panic(err)
			}
		}()
		jwtFile, _ := os.Open("templates/JWT.txt")
		jwtTemplate, _ := io.ReadAll(jwtFile)
		file.Write(jwtTemplate)
		exec.Command("go", "mod", "tidy").Run()
	},
}
