package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello from gin!")
	})

	router.Run(":5000")
}
