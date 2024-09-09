package server

import (
	"ltf/internal/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.LoadHTMLFiles("././web/template/index.html")
	router.LoadHTMLGlob("././web/template/*")

	router.Static("/static", "././web/static/")

	router.GET("/uploads", controller.GetUploadsHandler)

	router.GET("/files/*filePath", controller.GetFilesHandler)

	router.Run(":5000")
}
