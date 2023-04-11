package api

import (
	"airportal/db"
	"airportal/models"
	"log"

	"github.com/gin-gonic/gin"
)

func Cancellation(c *gin.Context) {
	var formData models.User

	if err1 := c.ShouldBind(&formData); err1 != nil {
		ResponseError(c, "参数错误", err1.Error())
		return
	}

	stmt, err := db.Db().Prepare("delete from users where username=?")
	if err != nil {
		log.Fatal(err)
	}

	_, err2 := stmt.Exec(formData.Username)

	if err2 != nil {
		ResponseError(c, "删除失败", err2)
		return
	}

	ResponseSuccess(c, "删除成功")

}
