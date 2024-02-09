package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func CreateJWT(cmd *cobra.Command, args []string) {
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
}

func CreateJWTMiddleware(cmd *cobra.Command, args []string) {
	var (
		file *os.File
		err  error
	)
	if len(args) == 0 {
		file, err = CreateJwtMiddlewareFile()
	} else {
		file, err = CreateJwtMiddlewareFile(args[0])
	}
	if err != nil {
		panic(err)
	}
	jwtFile, _ := os.Open("templates/jwt-middleware.txt")
	jwtTemplate, _ := io.ReadAll(jwtFile)
	file.Write(jwtTemplate)
	exec.Command("go", "mod", "tidy").Run()
	defer func() {
		file.Close()
		jwtFile.Close()
	}()
}

func CreateModel(cmd *cobra.Command, args []string) {
	var (
		file *os.File
		err  error
	)
	if len(args) == 0 {
		log.Fatalln("name is required")
	} else {
		file, err = CreateModelFile(args[0])
	}
	if err != nil {
		panic(err)
	}
	modelFile, _ := os.Open("templates/model.txt")
	modelTemplate, _ := io.ReadAll(modelFile)
	lines := strings.Split(string(modelTemplate), "\n")
	modelName := strings.Split(args[0], "/")[len(strings.Split(args[0], "/"))-1]
	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], "packageName", modelName)
		lines[i] = strings.ReplaceAll(lines[i], "ModelName", strings.ToUpper(modelName[:1])+strings.ToLower(modelName[1:]))
	}
	file.Write([]byte(strings.Join(lines, "\n")))
	exec.Command("go", "mod", "tidy").Run()
	defer func() {
		file.Close()
		modelFile.Close()
	}()
}

func CreateRepo(cmd *cobra.Command, args []string) {
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
		lines[i] = strings.ReplaceAll(lines[i], "Model", strings.ToUpper(repoName[:1])+strings.ToLower(repoName[1:]))
	}
	file.Write([]byte(strings.Join(lines, "\n")))
	exec.Command("go", "mod", "tidy").Run()
	defer func() {
		file.Close()
		repoFile.Close()
	}()
}
