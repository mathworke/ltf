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
	router.MaxMultipartMemory = 8 << 28 // 536 MiB

	router.GET("/", controller.GetIndedxPage)
	router.GET("/uploads", controller.GetUploadsFiles)
	router.GET("/uploads/:file", controller.GetFileHandler)
	router.POST("/uploads", controller.UploadFile)
	router.DELETE("/uploads/:filename", controller.DeleteFile)

	gin.SetMode(gin.ReleaseMode)

	router.Run(":5000")
}
