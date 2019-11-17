package service

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dybaseapi"
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

	endpoints.AddEndpointMapping("cn-hangzhou", "Dybaseapi", "dybaseapi.aliyuncs.com")

	// 创建client实例
	client, err := dybaseapi.NewClientWithAccessKey(
		"cn-hangzhou",           // 您的可用区ID
		"<AccessKeyId>",         // 您的Access Key ID
		"<AccessKeySecret>")     // 您的Access Key Secret
	if err != nil {
		// 异常处理
		panic(err)
	}

	queueName := "<QueueName>"
	messageType := "<MessageType>"

	var token *dybaseapi.MessageTokenDTO

	return true
}
