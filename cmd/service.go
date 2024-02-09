package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

var ServiceCmd = &cobra.Command{
	Use:   "service <path>",
	Short: "Create service",
	Run: func(cmd *cobra.Command, args []string) {
		CreateModel(cmd, args)
		CreateRepo(cmd, args)
		var (
			file *os.File
			err  error
		)
		if len(args) == 0 {
			log.Fatalln("name is required")
		} else {
			file, err = CreateServiceFile(args[0])
		}
		if err != nil {
			panic(err)
		}
		serviceFile, _ := os.Open("templates/service.txt")
		serviceTemplate, _ := io.ReadAll(serviceFile)
		lines := strings.Split(string(serviceTemplate), "\n")
		serviceName := strings.Split(file.Name(), "/")[len(strings.Split(file.Name(), "/"))-1]
		serviceName = strings.ReplaceAll(strings.TrimSuffix(serviceName, ".go"), "_", " ")
		serviceName = strings.Split(serviceName, " ")[0]
		for i := range lines {
			lines[i] = strings.ReplaceAll(lines[i], "packageName", serviceName)
			lines[i] = strings.ReplaceAll(lines[i], "service", serviceName)
			lines[i] = strings.ReplaceAll(lines[i], "Model", strings.ToUpper(serviceName[:1])+strings.ToLower(serviceName[1:]))
		}
		file.Write([]byte(strings.Join(lines, "\n")))
		exec.Command("go", "mod", "tidy").Run()
		defer func() {
			file.Close()
			serviceFile.Close()
		}()
	},
}
