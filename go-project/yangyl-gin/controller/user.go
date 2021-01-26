package controller

import (
	"yangyl-gin/gin"
)

func Login(ctx *gin.Context) {
	ctx.String(200, "hello")
}
