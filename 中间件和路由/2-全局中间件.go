package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
	Age  int
}

// 全局中间件
func m10(c *gin.Context) {
	fmt.Println("m10")
}

func m11(c *gin.Context) {
	fmt.Println("m11")
}

// Data 中间件传值
func Data(c *gin.Context) {
	c.Set("name", "cfd") //第二个参数可以为任意值

	c.Set("user", User{
		Name: "xzk",
		Age:  20,
	})
}

func main() {
	r := gin.Default()
	r.Use(m10, m11, Data) //此时任何一个请求发生的时候都会同时使用这谢个中间件,执行顺序为全局中间件，某个请求的中间件
	r.GET("/m10", func(c *gin.Context) {
		fmt.Println("index")
		c.JSON(200, gin.H{"msg": "成功"})
	})
	r.GET("/m11", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "成功"})
	})

	r.GET("data", func(c *gin.Context) {
		name, _ := c.Get("name") //第二个返回值为bool类型，判断是否存在这个值

		user, _ := c.Get("user")
		_user := user.(User) //当我们传递回来的是一个结构体并且我们想获取其中某一个具体的值，我们就需要断言

		c.JSON(200, gin.H{
			"msg":  name,
			"user": _user.Name,
		}) //可以让视图将中间件的数据获取下来
	})
	r.Run(":80")
}
