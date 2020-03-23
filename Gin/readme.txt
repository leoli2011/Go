gin  echo beego

gin
--------------------------
1）go的web框架
2)性能比较好
3）使用简单


gin helloworld
------------------------
1)确保已经安装go环境（goroot gopath配置）
2）go get gopkg.in/gin-goninc/gin.v1(需要翻墙)
   got get github.com/gin-gonic/gin.git (gin文件夹)

3)http://localhost:8080/

Helloworld

4)修改端口



注意点：

1）gin 需要翻墙
got get github.com/gin-gonic/gin.git

2）把课程对应的包拷贝到gopath对应的src和github.com


go路由
--------------------------------
httprouter库一个功能
gin的路由就是来自于httprouter库，因此httprouter具有的功能gin也具有

编码实战：
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()
	router.GET("/user/:name/*age", func(context *gin.Context) {
		name := context.Param("name")
		age := context.Param("age")
		message := name +" is "+age
		context.String(http.StatusOK,message)
	})
	router.Run()

}


注意点：
---------------
gin框架中除了使用：来代替变量，也可以使用*来代替，不过使用*代替表示为任意值


post 测试的俩个软件
--------------
1）postman

2）curl
1）下载https://curl.haxx.se/download.html
2）解压
3）进入对应的exe 【C:\it1804班级\go语言\gin\软件\curl-7.60.0\I386】
4）curl www.baidu.com

注册到环境里面
----------------
1）把exe文件拷贝到c:\windows\system32 文件中
2）配置系统环境变量
path=C:\it1804班级\go语言\gin\软件\curl-7.60.0\I386

gin post
-----------


curl模拟常见的场景
--------------
1)application/json (request中发送json数据用post方式发送Content-type用application/json)
2)application/x-www-form-urlencoded(常见的表达提交，把query string的内容 放到的body体内，同样需要urlencoded)
3）application/xml （http作为传输协议，xml作为编码方式远程调用规范）
4）multipart/form-data(图片上传)

代码实战 post请求
--------------------
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		test := c.DefaultPostForm("name","zhangsan")

		c.JSON(http.StatusOK,gin.H{"status":gin.H{"status_code":http.StatusOK,"status":
			"ok"},"message":message,"name":test})
	})


	router.Run()

}


请求方式
curl：
>curl -X POST http://127.0.0.1:8080/form_post -H "Content-Type:application/x-www-form-urlencoded" -d "message=hello&nama=zhangsan"

postman:
http://127.0.0.1:8080/form_post

{
  "message": "helloworld",
  "name": "wangwu",
  "status": {
    "status": "ok",
    "status_code": 200
  }
}






gin 文件上传
--------------------------------
post 方法 类型使用multipart/form-data 


代码实战
------------

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main()  {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		 
        name := c.PostForm("name") //根据name属性获取文件名
         fmt.Println(name)

       file,err :=  c.FormFile("upload")
       if err !=nil{
       	c.String(http.StatusBadRequest,"a bad request")
       	return
	   }

	   filename := file.Filename
	   fmt.Println("=============",filename)

	   if err := c.SaveUploadedFile(file,filename);err !=nil{
		   c.String(http.StatusBadRequest,"upload file err:%s",err.Error())
		   return
	   }

	   c.String(http.StatusCreated,"upload successful")

	})
	router.Run()
}


注意：
--------------------------------------------------------
name := c.PostForm("name") //根据name属性获取文件名

file,err :=  c.FormFile("upload") 获取文件
 filename := file.Filename 获取文件名

c.SaveUploadedFile 保存文件到服务器对应的路径

















