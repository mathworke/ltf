package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUploadsHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello from Gin!")
}
