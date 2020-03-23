package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	//router.GET("/user/:name/:age", func(context *gin.Context) {
	router.GET("/user/:name/*age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		message := name + " is " + age
		context.String(http.StatusOK, message)
	})
	router.Run()
}
