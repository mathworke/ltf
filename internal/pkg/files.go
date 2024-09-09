package pkg

import (
	"errors"
	"os"
	"runtime"
)

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

func GetFiles(path string) (files []string, err error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		files = append(files, file.Name())
	}

	return files, nil
}
