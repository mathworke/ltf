package controller

import (
	"ltf/internal/pkg"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetUploadsHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "upload.html", gin.H{})
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
			// ctx.JSON(http.StatusInternalServerError, gin.H{
			// "time":  time.Now(),
			// "error": err.Error(),
			// })
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"Files": files,
			})
			return
		}
		// ctx.JSON(http.StatusOK, gin.H{
		// "time":  time.Now(),
		// "files": files,
		// })
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Files": files,
		})
		return
	}

	ctx.File(string(path))

}
