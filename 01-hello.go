package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //创建一个默认路由引擎

	//发出一个GET请求  "/"为请求路径
	//当客户端以GET方法请求请求路径的时候，就会执行后面的匿名函数
	//绑定路由规则和路由函数，访问'/'的路由，将由对应的函数来处理
	r.GET("/", func(c *gin.Context) {

		//返回JSON格式的数据
		c.JSON(200, gin.H{
			"messgae": "Hello world!",
		})

		//第一个参数为状态码，第二个参数为要输出的数据，可以是map类型的数据，也可以是一个自义定的 struct 类型数据
	})

	//启动监听，默认在 0.0.0.0:8080 启动服务
	r.Run()

	//用原生http服务的方式启动，run函数本质上是对这个的封装
	http.ListenAndServe(":8080", r)
}
