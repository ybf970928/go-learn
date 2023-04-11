package api

import (
	"airportal/constants"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code int, data interface{}) bool {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  constants.GetMsg(code),
		"data": data,
	})
	return true
}

func ResponseSuccess(c *gin.Context, data interface{}) bool {
	c.JSON(http.StatusOK, gin.H{
		"code": constants.SUCCESS,
		"msg":  constants.GetMsg(constants.SUCCESS),
		"data": data,
	})
	return true
}

func ResponseError(c *gin.Context, msg string, data interface{}) bool {
	c.JSON(http.StatusOK, gin.H{
		"code": constants.ERROR,
		"msg":  msg,
		"data": data,
	})
	return true
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 安全性考虑，*是所有的域名都可以读取到数据，可能会读取到cookie等敏感数据
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Max-Age", "1800")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}
