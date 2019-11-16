package service

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	mnsDomain = "1943695596114318.mns.cn-hangzhou.aliyuncs.com"
)

/* 定义结构体 */
type SmsService struct {
}

//发送短信验证码
func (smsService SmsService) SendSmsCode(phone string) bool {
	code := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	strCode := string(code)
	return true
}
