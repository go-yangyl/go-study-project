package controller

import (
	"fmt"
	"yangyl-gin/gin"
)

var num = 10000

func Login(ctx *gin.Context) {
	num--
	fmt.Println(num)
}
