package api

import (
	"airportal/config"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
)

type Image struct {
	Path string `json:"path"`
}

type Images struct {
	Image
}

func ImageList(c *gin.Context) {

	res := make([]Images, 0)
	files, _ := os.ReadDir(config.Upload_url)

	var reg *regexp.Regexp

	for _, file := range files {
		reg, _ = regexp.Compile("jpeg|png|JPG|webp|gif|apk")
		matched := reg.MatchString(file.Name())
		if matched {
			res = append(res, Images{Image{Path: file.Name()}})
		}

	}
	data := make(map[string]interface{})
	data["files"] = res
	data["total"] = len(res)
	ResponseSuccess(c, data)
}
