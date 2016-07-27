package main

import (
	"./resources"
	"github.com/ahausamm/sc-helpers"
	"github.com/gin-gonic/gin"
	"fmt"
)

var router *gin.Engine

var GlobalKey = []byte("abcdefjdfhwuezfh")

func main() {
	var ciphertext, plaintext []byte
	var err error
	plaintext = []byte("text to encrypt")

	if ciphertext, err = encryption.Encrypt(GlobalKey,plaintext); err != nil {
		panic(err.Error())
	}
	fmt.Println(string(ciphertext[:]))

	if plaintext, err = encryption.Decrypt(GlobalKey,string(ciphertext[:])); err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n",plaintext)

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
