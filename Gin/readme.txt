gin  echo beego

gin
--------------------------
1��go��web���
2)���ܱȽϺ�
3��ʹ�ü�


gin helloworld
------------------------
1)ȷ���Ѿ���װgo������goroot gopath���ã�
2��go get gopkg.in/gin-goninc/gin.v1(��Ҫ��ǽ)
   got get github.com/gin-gonic/gin.git (gin�ļ���)

3)http://localhost:8080/

Helloworld

4)�޸Ķ˿�



ע��㣺

1��gin ��Ҫ��ǽ
got get github.com/gin-gonic/gin.git

2���ѿγ̶�Ӧ�İ�������gopath��Ӧ��src��github.com


go·��
--------------------------------
httprouter��һ������
gin��·�ɾ���������httprouter�⣬���httprouter���еĹ���ginҲ����

����ʵս��
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


ע��㣺
---------------
gin����г���ʹ�ã������������Ҳ����ʹ��*�����棬����ʹ��*�����ʾΪ����ֵ


post ���Ե��������
--------------
1��postman

2��curl
1������https://curl.haxx.se/download.html
2����ѹ
3�������Ӧ��exe ��C:\it1804�༶\go����\gin\���\curl-7.60.0\I386��
4��curl www.baidu.com

ע�ᵽ��������
----------------
1����exe�ļ�������c:\windows\system32 �ļ���
2������ϵͳ��������
path=C:\it1804�༶\go����\gin\���\curl-7.60.0\I386

gin post
-----------


curlģ�ⳣ���ĳ���
--------------
1)application/json (request�з���json������post��ʽ����Content-type��application/json)
2)application/x-www-form-urlencoded(�����ı���ύ����query string������ �ŵ���body���ڣ�ͬ����Ҫurlencoded)
3��application/xml ��http��Ϊ����Э�飬xml��Ϊ���뷽ʽԶ�̵��ù淶��
4��multipart/form-data(ͼƬ�ϴ�)

����ʵս post����
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


����ʽ
curl��
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






gin �ļ��ϴ�
--------------------------------
post ���� ����ʹ��multipart/form-data 


����ʵս
------------

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func main()  {
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		 
        name := c.PostForm("name") //����name���Ի�ȡ�ļ���
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


ע�⣺
--------------------------------------------------------
name := c.PostForm("name") //����name���Ի�ȡ�ļ���

file,err :=  c.FormFile("upload") ��ȡ�ļ�
 filename := file.Filename ��ȡ�ļ���

c.SaveUploadedFile �����ļ�����������Ӧ��·��

















