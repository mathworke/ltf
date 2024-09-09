package controller

import (
	"ltf/internal/pkg"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUploadsHandler(ctx *gin.Context) {
	// files, err := pkg.GetFiles("")
	// if err != nil {
	// ctx.String(http.StatusInternalServerError, "Internal Serverv Error")
	// }

	// for index, file := range files {
	// fmt.Printf("#%d: - %v\n", index, file)
	// }

	msg := "Hello from Gin!"

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "fraom controller",
		"Msg":   msg,
	})
}

func GetFilesHandler(ctx *gin.Context) {
	path := []byte(strings.Replace(ctx.Param("filePath"), "/", "\\", -1))

	rootPath, err := pkg.UserHomeDir()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	path = append([]byte(rootPath), path...)

	fileInfo, err := os.Stat(string(path))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if fileInfo.IsDir() {
		files, err := pkg.GetFiles(string(path))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"time":  time.Now(),
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"time":  time.Now(),
			"files": files,
		})
		return
	}

	ctx.File(string(path))

}
