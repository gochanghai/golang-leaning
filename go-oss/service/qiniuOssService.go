package service

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	_ "os"
	"strconv"
	_ "strconv"
	"strings"
	"time"
)

var (
	accessKey = "e0tgfo2PWicIByumNbofbvPFpJKe3aMAMNHvWnF1"
	secretKey = "eA3qQKKeyp3clgCFlYFCZq4cuSPUwDxnW_ayRF1-"
	bucket    = "2199-temp"
	path      = "q06tnjwyn.bkt.clouddn.com"
)

/* 定义结构体 */
type QiniuOSS struct {
}

// 自定义返回值结构体
type MyPutRet struct {
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	Fsize  int    `json:"size"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// 上传文件
func (qiniuOSS QiniuOSS) UploadFile(filePath string) {
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}

	sce := strings.Split(filePath, ".")
	key := strconv.FormatInt(time.Now().Unix(), 10) + "." + sce[1]
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": key,
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, filePath, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
}

// 下载文件
func (qiniuOSS QiniuOSS) DownloadFile(filePath string) string {
	var filename = ""
	return filename
}
