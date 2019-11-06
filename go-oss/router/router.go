package router

import (
	"github.com/gin-gonic/gin"
	. "golang-leaning/go-oss/apis"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	// aliyun
	router.POST("/uploadFile2Aliyun", UploadFile2Aliyun)
	router.POST("/download2Aliyun", DownloadFile2Aliyun)
	// qiniu
	router.POST("/uploadFile2Qiniu", UploadFile2Qiniu)
	router.POST("/download2Qiniu", DownloadFile2Qiniu)
	return router
}
