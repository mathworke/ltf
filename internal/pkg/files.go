package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
)

type Files struct {
	Name   string
	IsDir  bool
	Path   string
	Extend string
}

func UserHomeDir() (string, error) {
	env, enverr := "HOME", "$HOME"
	switch runtime.GOOS {
	case "windows":
		env, enverr = "USERPROFILE", "%userprofile%"
	case "plan9":
		env, enverr = "home", "$home"
	}
	if v := os.Getenv(env); v != "" {
		return v, nil
	}
	// On some geese the home directory is not always defined.
	switch runtime.GOOS {
	case "android":
		return "/sdcard", nil
	case "ios":
		return "/", nil
	}
	return "", errors.New(enverr + " is not defined")
}

func GetFiles() (files []string, err error) {
	HOME, err := UserHomeDir()
	if err != nil {
		log.Fatalln(err.Error())
	}
	rootPath := fmt.Sprintf("%s/root/", HOME)

	dir, err := os.ReadDir(rootPath)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		// partsFile := strings.Split(file.Name(), ".")
		// extend := partsFile[len(partsFile)-1]
		// files = append(files, Files{
		// 	Name:   file.Name(),
		// 	IsDir:  file.IsDir(),
		// 	Path:   file.Name(),
		// 	Extend: extend,
		// })
		files = append(files, file.Name())
	}

	return files, nil
}
