package cmd

import (
	"os"
	"os/exec"
	"strings"
)

func CreateFile(path, defaultPath string) *os.File {
	path = strings.TrimSuffix(strings.TrimSuffix(path, "/"), ".go")
	paths := strings.Split(path, "/")
	if len(paths) == 1 {
		exec.Command("mkdir", defaultPath).Run()
		path = defaultPath + "/" + path
	} else {
		exec.Command("mkdir", "-p", strings.Join(paths[:len(paths)-1], "/")).Run()
	}
	file, err := os.Create(path + ".go")
	if err != nil {
		return nil
	}
	return file
}

func CreateJwtFile(name ...string) (*os.File, error) {
	path := "auth"
	if len(name) != 0 {
		path = name[0]
	}
	file := CreateFile(path, "auth")
	if file == nil {
		return nil, os.ErrNotExist
	}
	return file, nil
}

func CreateRepoFile(name string) (*os.File, error) {
	fileName := strings.Split(name, "/")[len(strings.Split(name, "/"))-1]
	name += "/" + fileName
	if _, err := os.Open(name + ".go"); err == nil {
		name += "_repo"
	}
	file := CreateFile(name, name)
	if file == nil {
		return nil, os.ErrNotExist
	}
	return file, nil
}

func CreateModelFile(name string) (*os.File, error) {
	fileName := strings.Split(name, "/")[len(strings.Split(name, "/"))-1]
	name += "/" + fileName
	file := CreateFile(name, name)
	if file == nil {
		return nil, os.ErrNotExist
	}
	return file, nil
}

func CreateServiceFile(name string) (*os.File, error) {
	fileName := strings.Split(name, "/")[len(strings.Split(name, "/"))-1]
	name += "/" + fileName
	if _, err := os.Open(name + ".go"); err == nil {
		name += "_service"
	}
	file := CreateFile(name, name)
	if file == nil {
		return nil, os.ErrNotExist
	}
	return file, nil
}
