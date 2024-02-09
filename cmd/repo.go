package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var RepoCmd = &cobra.Command{
	Use:   "repo <path>",
	Short: "Create repository",
	Run: func(cmd *cobra.Command, args []string) {
		CreateModel(cmd, args)
		var (
			file *os.File
			err  error
		)
		if len(args) == 0 {
			log.Fatalln("name is required")
		} else {
			file, err = CreateRepoFile(args[0])
		}
		if err != nil {
			panic(err)
		}
		repoFile, _ := os.Open("templates/repo.txt")
		repoTemplate, _ := io.ReadAll(repoFile)
		lines := strings.Split(string(repoTemplate), "\n")
		repoName := strings.Split(file.Name(), "/")[len(strings.Split(file.Name(), "/"))-1]
		repoName = strings.ReplaceAll(strings.TrimSuffix(repoName, ".go"), "_", " ")
		repoName = strings.Split(repoName, " ")[0]
		for i := range lines {
			lines[i] = strings.ReplaceAll(lines[i], "packageName", repoName)
			lines[i] = strings.ReplaceAll(lines[i], "repository", repoName)
			lines[i] = strings.ReplaceAll(lines[i], "Repository", strings.ToUpper(repoName[:1])+strings.ToLower(repoName[1:]))
		}
		file.Write([]byte(strings.Join(lines, "\n")))
		exec.Command("go", "mod", "tidy").Run()
		defer func() {
			file.Close()
			repoFile.Close()
		}()
	},
}
