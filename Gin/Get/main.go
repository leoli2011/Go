package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/user/", func(context *gin.Context) {
		firstname := context.DefaultQuery("firstname", "Li")
		lastname := context.Query("lastname")
		message := firstname + " and " + lastname
		context.String(http.StatusOK, message)
	})
	router.Run()
}
