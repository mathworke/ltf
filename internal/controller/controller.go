package controller

import (
	"fmt"
	"io"
	"ltf/internal/pkg"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUploadsFiles(ctx *gin.Context) {
	files, err := pkg.GetFiles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"time":  time.Now(),
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, files)
}

func GetIndedxPage(ctx *gin.Context) {
	files, err := pkg.GetFiles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"time":  time.Now(),
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Uploads": files,
	})
}

func GetFileHandler(ctx *gin.Context) {
	filepath := fmt.Sprintf("%v%v", pkg.Rootpath(), ctx.Param("file"))

	ctx.File(filepath)
}

func UploadFile(ctx *gin.Context) {
	file, err := ctx.FormFile("files")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"time":  time.Now(),
			"error": err.Error(),
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"time":  time.Now(),
			"error": err.Error(),
		})
		return
	}
	defer src.Close()

	dst, err := os.Create(filepath.Join(pkg.Rootpath(), file.Filename))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"time":  time.Now(),
			"error": err.Error(),
		})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"time":  time.Now(),
			"error": err.Error(),
		})
		return
	}

	ctx.Redirect(http.StatusFound, "/")
}

func DeleteFile(ctx *gin.Context) {
	filename := ctx.Param("filename")

	if err := os.Remove(path.Join(pkg.Rootpath(), filename)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"time":  time.Now(),
			"error": err.Error(),
		})
		return
	}

	files, err := pkg.GetFiles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"time":  time.Now(),
			"error": err.Error(),
		})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Uploads": files,
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
		if err != nil {
			ctx.HTML(http.StatusOK, "index.html", gin.H{})
			return
		}

		ctx.HTML(http.StatusOK, "index.html", gin.H{})
		return
	}

	ctx.File(string(path))
}
