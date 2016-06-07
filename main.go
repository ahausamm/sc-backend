package main

import (
	"./resources"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("/tokens/create/:InstanceId/:EncryptedString", tokens.CreateToken)
	}

	router.Run(":9090")
	return
}
