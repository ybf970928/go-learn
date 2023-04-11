package api

import (
	"airportal/config"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {

	srcFile, head, _ := c.Request.FormFile("file")
	fmt.Printf("head.Filename: %v\n", head.Filename)
	// 上传文件至指定的完整文件路径
	dstFile, err := os.Create(config.Upload_url + head.Filename)
	if err != nil {
		ResponseError(c, "创建文件夹失败", err.Error())
		return
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		ResponseError(c, "上传失败", err.Error())
		return
	}

	defer dstFile.Close()
	ResponseSuccess(c, "上传成功")
}
