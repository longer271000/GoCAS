package main

import "github.com/gin-gonic/gin"

func main()  {

	r:=gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"Blog":"www.flysnow.org",
			"wechat":"flysnow_org",
			"host":"localhost",
		})
	})
	r.Run(":8080")

}

