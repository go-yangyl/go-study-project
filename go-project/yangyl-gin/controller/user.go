package controller

import (
	"fmt"
	"yangyl-gin/gin"
)

func Login(ctx *gin.Context) {
	fmt.Println("控制器消息")
	ctx.String(200, "hello")
}
