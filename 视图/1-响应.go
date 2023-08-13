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

func Html(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"username": "cfd"}) //可以将参数传递给html，在index.html页面也有部分细节
}

// Redirect 301是临时重定向，302是永久重定向：https://blog.csdn.net/qq_43968080/article/details/107355758
func Redirect(c *gin.Context) {
	c.Redirect(301, "https://cn.bing.com/?mkt=zh-CN")
}

func main() {
	r := gin.Default()

	r.GET("/string", String) //响应字符串

	r.GET("/json1", Json1) //响应json
	r.GET("/json2", Json2) //响应json
	r.GET("/json3", Json3) //响应json

	r.GET("/xml", Xml) //响应xml

	r.GET("/yaml", Yaml) //响应yaml

	r.LoadHTMLGlob("templates/*") //加载模板目录下的所有模板文件
	r.GET("/html", Html)

	//加载静态资源，将文件加载到网页上，前一个参数表示网页请求的路由，后一个参数表示文件存储在哪的路径
	//这样就可以选择将哪些文件加载到网页，因为不可能一次将整个项目全部加载到网页，这样的话就很危险
	//在goland中，没有相对文件路径，只有相对的项目路径。也就是说，寻找文件的时候是从GinStudy开始往下找，所以文件路径可以直接写成 "static/Doraemon.jpg"
	r.StaticFile("/哆啦a梦", "static/Doraemon.jpg")

	//网页请求这个静态网页的前缀，第二个参数是请求的目录
	//像这样，就只能访问到 hello.txt，不能访问到 world.txt
	r.StaticFS("/static", http.Dir("static/static"))
	//上述两个请求的第一个参数不能有重复的前缀

	r.GET("/bing", Redirect) //重定向

	r.Run(":80") //端口为80表示访问 127.0.0.1 就能成功，不需要加上端口号
}
