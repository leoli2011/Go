package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func main() {
	router := gin.Default()

	router.POST("/upload", func(context *gin.Context) {

		name := context.PostForm("name")  //根据name属性获取文件名
		fmt.Println("name=", name)

		file, err := context.FormFile("upload")
		if err != nil {
			context.String(http.StatusBadRequest, "a bad request")
			return
		}

		filename := file.Filename
		fmt.Println("================", filename)

		if err := context.SaveUploadedFile(file, name); err != nil {
			context.String(http.StatusBadRequest, "upload file err:%s", err.Error())
			return
		}
		context.String(http.StatusOK, "upload successful")
	})

	router.Run()
}