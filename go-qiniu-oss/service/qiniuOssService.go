package service

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	_ "os"
	_ "strconv"
	"time"
)

var (
	accessKey = "e0tgfo2PWicIByumNbofbvPFpJKe3aMAMNHvWnF1"
	secretKey = "eA3qQKKeyp3clgCFlYFCZq4cuSPUwDxnW_ayRF1-"
	bucket = "2199-temp"
	path = "q06tnjwyn.bkt.clouddn.com"
)

// 自定义返回值结构体
type MyPutRet struct {
	Key    string `json:"key"`
	Hash   string `json:"hash"`
	Fsize  int `json:"size"`
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

// 上传文件
func uploadFile(file string)  {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}

	key := time.Now().UnixNano()
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := MyPutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, string(key), file, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)
}