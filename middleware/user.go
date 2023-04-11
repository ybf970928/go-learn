package middleware

import (
	"airportal/api"
	"airportal/constants"
	"airportal/library"
	"time"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {

	var data interface{}
	token := c.GetHeader("token")

	if token == "" {
		api.ResponseError(c, constants.GetMsg(constants.ErrorAuthCheckTokenFail), data)
		c.Abort()
		return
	}

	claims, err := library.ParseToken(token)

	if err != nil {
		api.ResponseError(c, "token失效", data)
		c.Abort()
		return
	}

	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		api.ResponseError(c, "token过期", data)
		c.Abort()
		return
	}

	c.Next()
}
