package api

import (
	"airportal/db"
	"airportal/library"
	"airportal/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var formData models.User

	if err1 := c.ShouldBind(&formData); err1 != nil {
		ResponseError(c, "参数错误", err1.Error())
		return
	}

	var user models.User

	err := db.Db().QueryRow("select id, username, password from users where username=?", formData.Username).Scan(&user.ID, &user.Username, &user.Password)

	if err != nil {
		ResponseError(c, "账号不存在", nil)
		return
	}

	if user.Password != formData.Password {
		ResponseError(c, "密码错误", "")
		return
	}

	data := make(map[string]interface{})

	toekn, err1 := library.GenerateToken(user.ID)

	if err1 != nil {
		ResponseError(c, "token生成失败", "")
		return
	}

	data["token"] = toekn

	data["id"] = user.ID

	data["username"] = user.Username

	ResponseSuccess(c, data)

}
