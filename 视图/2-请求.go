package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

// Query 当浏览器上出现像 ?user=zhangsan 这样的东西的时候，这个函数能够获取到这些参数
func Query(c *gin.Context) {
	fmt.Println(c.Query("user"))               //能够拿到查询参数
	fmt.Println(c.GetQuery("user"))            //能够判断是否有这个参数，同时拿到查询参数
	fmt.Println(c.QueryArray("user"))          //能够拿到多个相同地查询参数
	fmt.Println(c.DefaultQuery("addr", "湖北省")) //如果用户没有传递参数，则使用默认值湖北省
}

// Param 当参数为动态改变的时候，可以使用这个来获取参数
// 请求路径：param/任意数字或字符，都能够获取到这个数字和字符
func Param(c *gin.Context) {
	fmt.Println(c.Param("user_id"))
}

// Form 能够接受两种：form-data、x-www-form-urlencoded 格式
func Form(c *gin.Context) {
	fmt.Println(c.PostForm("name"))
	fmt.Println(c.PostFormArray("name"))
	fmt.Println(c.DefaultPostForm("addr", "湖北省")) //如果用户没有传递参数，则使用默认值湖北省
	fmt.Println(c.MultipartForm())                //接收所有的form参数包括文件
}

// Raw 解析原始参数
func Raw(c *gin.Context) {
	body, _ := c.GetRawData()
	ContentType := c.GetHeader("Content-Type") //请求头

	//当需要解析json数据的时候,这样子就只有json数据能够被输出，其他数据不会执行输出
	switch ContentType {
	case "application/json":
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := json.Unmarshal(body, &user) //将json数据解析成结构体储存在user中
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(user)
	}

	//fmt.Println(string(body))
}

func main() {
	r := gin.Default()

	r.GET("/query", Query) //查询参数

	r.GET("/param/:user_id", Param) //动态参数

	r.POST("/form", Form) //表单参数

	r.POST("/raw", Raw) //原始参数

	r.Run(":80")
}
