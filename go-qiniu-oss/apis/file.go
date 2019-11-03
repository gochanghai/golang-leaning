package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"golang-leaning/go-qiniu-oss/service"
)
/**上传方法**/
func UploadFile(c *gin.Context){
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
	//创建文件
	out, err := os.Create("static/uploadfile/"+filename)
	//注意此处的 static/uploadfile/ 不是/static/uploadfile/
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// 把文件上传到指定目录
	//_, err = io.Copy(out, file)
	service.AliyunOSS.UploadFile()
	if err != nil {
		log.Fatal(err)
	}
	c.String(http.StatusCreated, "upload successful")
}
