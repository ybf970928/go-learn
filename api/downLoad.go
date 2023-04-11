package api

import (
	"airportal/config"

	"github.com/gin-gonic/gin"
)

func DownLoadImage(c *gin.Context) {
	path := c.Param("path")
	dst := config.Upload_url + path

	c.File(dst)
}
