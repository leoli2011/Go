package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func loginEndpoint(c *gin.Context)  {
	fmt.Println("this is a login method")
}

func submitEndpoint(c *gin.Context)  {
	fmt.Println("this is a submit method")
}
func readEndpoint(c *gin.Context)  {
	fmt.Println("this is a read method")
}
func main() {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		v1.GET("login", loginEndpoint)
		v1.GET("submit", submitEndpoint)
		v1.GET("read", readEndpoint)
	}

	v2 := router.Group("v2")
	{
		v2.GET("login", loginEndpoint)
		v2.GET("submit", submitEndpoint)
		v2.GET("read", readEndpoint)
	}
	router.Run()
}
