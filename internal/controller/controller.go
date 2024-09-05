package controller

import (
	"fmt"
	"ltf/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUploadsHandler(ctx *gin.Context) {
	files, err := pkg.GetUploadsFiles()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Serverv Error")
	}

	for index, file := range files {
		fmt.Printf("#%d: - %v\n", index, file.Name())
	}

	msg := "Hello from Gin!"

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "fraom controller",
		"Msg":   msg,
	})
}
