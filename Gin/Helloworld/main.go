package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "HelloWorld")
	})

	router.Run(":5555")
}
