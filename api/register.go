package api

import (
	"airportal/db"
	"airportal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var formData models.User

	if err1 := c.ShouldBind(&formData); err1 != nil {
		ResponseError(c, "参数错误", err1.Error())
		return
	}

	stmt, err := db.Db().Prepare("insert into users(username, password) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := stmt.Exec(formData.Username, formData.Password)
	if err2 != nil {
		ResponseError(c, "注册失败", err2)
		return
	}
	ResponseSuccess(c, "注册成功")
}
