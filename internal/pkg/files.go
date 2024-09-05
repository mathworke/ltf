package pkg

import (
	"io/fs"
	"os"
)

func GetUploadsFiles() ([]fs.DirEntry, error) {
	return os.ReadDir("././cmd/ltf/upload")
}
