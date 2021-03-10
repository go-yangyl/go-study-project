package main

import (
	"yangyl-gin/controller"
	"yangyl-gin/gin"
)

func main() {
	g := gin.New()

	admin := g.Group("/admin")
	admin.Use(func(c *gin.Context) {
		c.Next()
	})

	admin.GET("/user", controller.Login)

	g.Run(":8080")
}
