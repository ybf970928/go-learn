package api

import (
	"airportal/db"
	"airportal/models"

	"github.com/gin-gonic/gin"
)

func Coffees(c *gin.Context) {

	// page := c.Query("page")
	size := c.Query("size")

	data := make([]models.Coffee, 0)

	rows, err := db.Db().Query("select name, brand, recommendations from coffees limit ?,?", size)
	if err != nil {
		ResponseError(c, "查询出错", nil)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var coffee models.Coffee
		err := rows.Scan(&coffee.Name, &coffee.Brand, &coffee.Recommendations)
		if err != nil {
			ResponseError(c, "查询出错", nil)
			return
		}
		data = append(data, coffee)
	}

	ResponseSuccess(c, data)
}
