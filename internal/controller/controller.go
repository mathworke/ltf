package controller

import (
	"fmt"
	"ltf/internal/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUploadsHandler(ctx *gin.Context) {
	files, err := pkg.GetUploadsFiles("")
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Internal Serverv Error")
	}

	for index, file := range files {
		fmt.Printf("#%d: - %v\n", index, file)
	}

	msg := "Hello from Gin!"

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "fraom controller",
		"Msg":   msg,
	})
}

func GetFiles(ctx *gin.Context) {
	path := ctx.Request.URL.Path

	path = strings.Replace(path, "/explorer/", "", 1)

	ctx.File(fmt.Sprintf("././root/%s", path))

	ctx.JSON(http.StatusOK, gin.H{
		"debug": string(path),
	})
}
