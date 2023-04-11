package router

import (
	"airportal/api"
	"airportal/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"http://127.0.0.1"})
	r.Use(api.CORSMiddleware())

	r.MaxMultipartMemory = 8 << 20

	r.POST("/login", api.Login)
	r.POST("/register", api.Register)
	r.POST("/cancellation", api.Cancellation)
	r.POST("/setpassword", api.Setpassword)

	fileGroup := r.Group("/images").Use(middleware.User)
	{
		fileGroup.GET("/getAllImages", api.ImageList)
		fileGroup.GET("/DownLoadImage/:path", api.DownLoadImage)
		fileGroup.POST("/upload", api.Upload)
	}

	coffeeGroup := r.Group("/coffees")
	{
		coffeeGroup.GET("/getCoffees", api.Coffees)
	}

	r.Run(":9527") // listen and serve on 0.0.0.0:9527

}
