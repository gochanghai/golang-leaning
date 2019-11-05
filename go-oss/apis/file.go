package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-leaning/go-oss/service"
	"io"
	"log"
	"net/http"
	"os"
)

var aliyunOSS service.AliyunOSS

/** 上传方法 **/
func UploadFile(c *gin.Context) {
	//得到上传的文件
	//file这个是uplaodify参数定义中的   'fileObjName':'file'
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename

	fmt.Println(file, err, filename)
	filePath := filename
	//创建文件
	out, err := os.Create(filePath)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// 把文件上传到指定目录
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	aliyunOSS.UploadFile(filePath)
	c.String(http.StatusCreated, "upload successful")
}

/** 下载文件 */
func DownloadFile(c *gin.Context) {
	filePath := c.PostForm("filePath")
	downloadedFile := aliyunOSS.DownloadFile(filePath)
	//对下载的文件重命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", downloadedFile))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(downloadedFile)
}
