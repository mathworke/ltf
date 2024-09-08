package server

import (
	"ltf/internal/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.LoadHTMLFiles("././web/template/index.html")

	router.Static("/static", "././web/static/")

	router.GET("/uploads", controller.GetUploadsHandler)

	router.GET("/explorer/*url", controller.GetFiles)

	router.Run(":5000")
}
