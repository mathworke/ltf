package main

import (
	"fmt"
	"log"
	"ltf/internal/pkg"
	"ltf/internal/server"
	"os"
)

func main() {

	{
		HOME, err := pkg.UserHomeDir()
		if err != nil {
			log.Fatalln(err.Error())
		}
		rootPath := fmt.Sprintf("%s/root/", HOME)
		if _, err := os.Stat(rootPath); err != nil {
			if os.IsNotExist(err) {
				err = os.Mkdir(rootPath, 0777)
				if err != nil {
					log.Fatalln(err.Error())
				}
			}
		}
	}

	server.Run()
}
