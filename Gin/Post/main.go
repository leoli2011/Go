package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/form_post", func(context *gin.Context) {
		message := context.PostForm("message")
		test := context.DefaultPostForm("name", "zhangsan")
		context.JSON(http.StatusOK,
			gin.H{"status": gin.H{"status_code":http.StatusOK,"status":"OK"},
			"message":message,
			"name":test})
	})
	router.Run()
}
