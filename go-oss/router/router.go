package router

import (
	"github.com/gin-gonic/gin"
	. "golang-leaning/go-oss/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/uploadFile", UploadFile)
	router.POST("/download", DownloadFile)
	return router
}
