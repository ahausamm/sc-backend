package main

import (
	"./resources"
	"github.com/gin-gonic/gin"
	"fmt"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/token/create", token.CreateToken)
	}

	fmt.Println("Management is ready")
	router.Run(":9090")
	return
}
