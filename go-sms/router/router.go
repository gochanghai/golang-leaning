package router

import (
	"github.com/gin-gonic/gin"
	. "golang-leaning/go-sms/controller"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/sendSms", SendSms)
	router.POST("/sendBatchSms", SendBatchSms)
	router.POST("/sendSmsCode", SendSmsCode)
	return router
}
