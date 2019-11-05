package service

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
	"os"
	"strings"
	"time"
)

/* 定义结构体 */
type AliyunOSS struct {
	radius float64
}

var (
	// Endpoint以深圳为例，其它Region请按实际情况填写。
	endpoint = "oss-cn-shenzhen.aliyuncs.com"
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId     = "LTAI4FhUxGHTWZ3XE7bcF2WS"
	accessKeySecret = "bsha9dwnJstxd0Ga2l6t8izADfO8Jp"
	bucketName      = "2199-temp"
	// <yourObjectName>上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg。
	objectName = ""
	// <yourLocalFileName>由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt。
	localFileName = ""
)

func handleError(err error) {
	log.Fatalln("Error:", err)
	os.Exit(-1)
}

// 上传文件
func (aliyunOSS AliyunOSS) UploadFile(filePath string) {
	sce := strings.Split(filePath, ".")
	objectName = string(time.Now().UnixNano()) + "." + sce[1]
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 读取本地文件。
	fd, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	defer fd.Close()
	// 上传文件。
	err = bucket.PutObject(objectName, fd)
	if err != nil {
		handleError(err)
	}
	log.Println("文件上传成功")
}

// 下载文件
func downloadedFile() {
	var downloadedFileName = ""
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 下载文件。
	err = bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {
		handleError(err)
	}
}
