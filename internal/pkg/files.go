package pkg

import (
	"os"
)

func GetUploadsFiles(path string) (files []string, err error) {
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		files = append(files, file.Name())
	}
	return files, nil
}
