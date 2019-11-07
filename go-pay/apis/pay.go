package apis

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-leaning/go-oss/service"
	"io"
	"log"
	"net/http"
	"os"
)

var aliyunOSS service.AliyunOSS

var qiniuOSS service.QiniuOSS

/** Qiniu上传方法 **/
func UploadFile2Qiniu(c *gin.Context) {
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
	qiniuOSS.UploadFile(filePath)
	c.String(http.StatusCreated, "upload successful")
}

/** Qiniu下载文件 */
func DownloadFile2Qiniu(c *gin.Context) {
	filePath := c.PostForm("filePath")
	download(filePath, "qiniu", c)
}

/** Aliyun上传方法 **/
func UploadFile2Aliyun(c *gin.Context) {
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

/** Aliyun下载文件 */
func PostDownloadFile2Aliyun(c *gin.Context) {
	filePath := c.PostForm("filePath")
	download(filePath, "aliyun", c)
}

/** Aliyun下载文件 */
func GetDownloadFile2Aliyun(c *gin.Context) {
	filePath := c.Query("filePath")
	download(filePath, "aliyun", c)
}

// 下载文件
func download(filePath string, ossType string, c *gin.Context) {
	var filename = ""
	if "aliyun" == ossType {
		filename = aliyunOSS.DownloadFile(filePath)
	} else {
		filename = qiniuOSS.DownloadFile(filePath)
	}
	//对下载的文件重命名
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	if !checkFileIsExist(filename) {
		file, err := os.Create(filename) //创建文件
		if err != nil {
			c.String(400, err.Error())
			return
		}
		buf := bufio.NewWriter(file) //创建新的 Writer 对象
		buf.WriteString("test")
		buf.Flush()
		defer file.Close()
	}
	//返回文件流
	c.File(filename)
	// 删除临时文件
	derr := os.Remove(filename)
	if derr != nil {
		// 删除失败
		log.Println("文件删除失败")
	} else {
		// 删除成功
		log.Println("文件删除成功")
	}
}

//判断文件是否存在  存在返回 true 不存在返回false
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
