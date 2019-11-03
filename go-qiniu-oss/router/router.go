package router

import (
	"github.com/gin-gonic/gin"
	. "golang-leaning/go-qiniu-oss/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/uploadFile", UploadFile)

	return router
}