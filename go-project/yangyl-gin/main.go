package main

import (
	"fmt"
	"yangyl-gin/controller"
	"yangyl-gin/gin"
)

func main() {
	g := gin.New()

	admin := g.Group("/admin")
	admin.Use(func(c *gin.Context) {
		fmt.Println("中间件消息")
		c.Next()
	})

	admin.GET("/user", controller.Login)

	g.Run(":8080")
}
