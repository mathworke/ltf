package main

import (
	"log"
	"ltf/internal/pkg"
	"ltf/internal/server"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	{
		rootPath := pkg.Rootpath()
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
