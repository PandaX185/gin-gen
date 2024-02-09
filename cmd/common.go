package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

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
