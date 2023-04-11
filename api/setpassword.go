package api

import (
	"airportal/db"
	"airportal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Setpassword(c *gin.Context) {
	var formData models.User

	if err := c.ShouldBind(&formData); err != nil {
		ResponseError(c, "参数错误", err.Error())
		return
	}

	var user models.User

	err1 := db.Db().QueryRow("select username, password from users where username=?", formData.Username).Scan(&user.Username, &user.Password)

	if err1 != nil {
		ResponseError(c, "账号不存在", nil)
		return
	}

	if user.Password == formData.Password {
		ResponseError(c, "新旧密码重复", nil)
		return
	}

	stmt, err2 := db.Db().Prepare("update users set password=? where username=?")
	if err2 != nil {
		log.Fatal(err2)
	}

	_, err3 := stmt.Exec(formData.Password, formData.Username)

	if err2 != nil {
		ResponseError(c, "密码设置错误", err3.Error())
		return
	}

	ResponseSuccess(c, "密码设置成功")
}
