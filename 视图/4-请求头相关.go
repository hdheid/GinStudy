package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	r := gin.Default()

	//请求头的获取方式
	r.GET("/", func(c *gin.Context) {
		//首字母的大小写不分，单词于单词间使用 - 连接
		//获取其中某一个请求头
		fmt.Println(c.GetHeader("User-Agent"))
		fmt.Println(c.GetHeader("user-Agent"))
		fmt.Println(c.GetHeader("user-agent"))
		//输出结果一样

		//Header是一个map数据结构，他的get方法和上述函数类似，只不过首字母必须大写
		//获取全部请求头
		fmt.Println(c.Request.Header)

		//自定义请求头，也不区分大小写
		fmt.Println(c.GetHeader("Token"))

		c.JSON(200, gin.H{
			"message": "成功",
		})
	})

	//爬虫和用户的区别
	r.GET("/python", func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")

		//使用正则表达式匹配

		//使用字符串包含匹配
		//表示只要字符串 userAgent 中出现了 python 字符串，就返回真
		if strings.Contains(userAgent, "python") {
			c.JSON(0, gin.H{"data": "是爬虫"})
			return
		}

		c.JSON(0, gin.H{"data": "不是爬虫"})

	})

	//设置响应头
	r.GET("/res", func(c *gin.Context) {
		//第一个参数是 key，第二个参数是 value
		c.Header("Token", "cfddfc")

		//这样设置之后，就不会是相应的JSON数据而是text数据，可以这样控制浏览器的行为
		c.Header("Content-Type", "application/text; charset=utf-8")
		c.JSON(200, gin.H{
			"message": "成功了",
		})
	})

	r.Run(":80")
}
