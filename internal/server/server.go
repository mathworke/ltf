package server

import (
	"ltf/internal/controller"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.GET("/uploads", controller.GetUploadsHandler)

	router.Run(":5000")
}
