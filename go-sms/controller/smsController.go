package controller

import (
	"github.com/gin-gonic/gin"
	"golang-leaning/go-sms/service"
	"net/http"
	"strconv"
)

var smsService service.SmsService

//发短信
func SendSms(c *gin.Context) {
	//user.Username = c.Request.FormValue("username")
	//user.Password = c.Request.FormValue("password")

}

//批量发送短信
func SendBatchSms(c *gin.Context) {
	//user.Username = c.Request.FormValue("username")
	//user.Password = c.Request.FormValue("password")

}

//发送短信验证码
func SendSmsCode(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	phone := c.Request.FormValue("phone")
	result := smsService.SendSmsCode(phone)
	if !result {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "发送失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "发送成功",
	})
}
