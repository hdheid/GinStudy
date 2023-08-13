package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func String(c *gin.Context) {
	c.String(200, "你好") //状态码可以修改，200表示正常响应
}

func Json1(c *gin.Context) {
	//json响应结构体
	type UserInfo struct {
		UserName string `json:"user_name"`
		Age      int    `json:"age"`

		Password string `json:"-"`
	} //定义一个结构体

	//如果有像密码一样的不能直接展示的数据，可以使用 `json:"-"` 让浏览器不进行渲染展示
	//加上 `json:"user_name"` 表示相应为JSON数据的时候，名字会改成这个，而不是结构体里面的那样

	user := UserInfo{
		UserName: "刘浩",
		Age:      19,
		Password: "123456",
	} //创建一个实例

	c.JSON(200, user) //响应这样一个JSON
}

func Json2(c *gin.Context) {
	//json响应map响应
	UserMap := map[string]string{
		"user_name": "刘浩",
		"age":       "19",
	}

	c.JSON(200, UserMap) //响应这样一个JSON
}

func Json3(c *gin.Context) {
	//直接响应JSON
	c.JSON(200, gin.H{
		"user-name": "袖子开",
		"age":       20,
	}) //使用gin的H可以直接写json文本进行响应
}

func Xml(c *gin.Context) {
	c.XML(200, gin.H{
		"user":    "hanru",
		"message": "hey",
		"status":  http.StatusOK,
		"data": gin.H{
			"year":  2023,
			"month": 8,
			"day":   13,
		},
	})
} //无论是json还是xml都是可以嵌套的,就像 data 这样

func Yaml(c *gin.Context) {
	c.YAML(200, gin.H{
		"user":    "hanru",
		"message": "hey",
		"status":  http.StatusOK,
		"data": gin.H{
			"year":  2023,
			"month": 8,
			"day":   13,
		},
	})
}

func main() {
	r := gin.Default()

	r.GET("/string", String) //响应字符串

	r.GET("/json1", Json1) //响应json
	r.GET("/json2", Json2) //响应json
	r.GET("/json3", Json3) //响应json

	r.GET("/xml", Xml) //响应xml

	r.GET("/yaml", Yaml) //响应yaml

	r.Run(":80") //端口为80表示访问 127.0.0.1 就能成功，不需要加上端口号
}
