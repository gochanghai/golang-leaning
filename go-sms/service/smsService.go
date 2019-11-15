package service

const (
	mnsDomain = "1943695596114318.mns.cn-hangzhou.aliyuncs.com"
)

/* 定义结构体 */
type SmsService struct {
}

//发送短信验证码
func (smsService SmsService) SendSmsCode(phone string) bool {

	return true
}
