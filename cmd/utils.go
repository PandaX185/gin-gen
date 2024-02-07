package cmd

import (
	"os"
	"os/exec"
	"strings"
)

func CreateFile(name ...string) (*os.File, error) {
	path := "login"
	if len(name) != 0 {
		path = name[0]
	}
	path = strings.TrimSuffix(strings.TrimSuffix(path, "/"), ".go")
	paths := strings.Split(path, "/")
	for i := 0; i < len(paths)-1; i++ {
		exec.Command("mkdir", paths[i]).Run()
	}
	if len(paths) == 1 {
		exec.Command("mkdir", "auth").Run()
		path = "auth/" + path
	}
	file, err := os.Create(path + ".go")
	if err != nil {
		return nil, err
	}
	return file, nil
}
