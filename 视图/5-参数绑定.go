package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name" form:"name" uri:"name"`
	Age  int    `json:"age" form:"age" uri:"age"`
	Sex  string `json:"sex" form:"sex" uri:"sex"`
}

// 如果获取到的body体里面的json数据和上述结构体的数据类型不一样，就会在这里出现报错，完成了json数据的绑定和校验
func _Should(c *gin.Context) {
	var userInfo UserInfo
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(200, gin.H{"msg": "出错了！"})
		return
	}

	c.JSON(200, userInfo)
}

// 绑定查询参数,在 params 中，使用 ShouldBindQuery 函数时，需要在tag后面加上 form
func _Query(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBindQuery(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "出错了！"})
		return
	}

	c.JSON(200, userInfo)
}

// tag 对应为 uri，绑定动态参数
func _Uri(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBindUri(&userInfo)
	if err != nil {
		c.JSON(200, gin.H{"msg": "出错了！"})
		return
	}

	c.JSON(200, userInfo)
}

// 可以绑定和校验 form-data 参数，tag 使用 form
func _Form(c *gin.Context) {
	var userInfo UserInfo
	err := c.ShouldBind(&userInfo) //这个函数会根据请求头自动地判断是哪一种数据
	if err != nil {
		c.JSON(200, gin.H{"msg": "出错了！"})
		return
	}

	c.JSON(200, userInfo)
}

func main() {
	r := gin.Default()

	r.POST("/", _Should)

	r.POST("/query", _Query)

	r.POST("/uri/:name/:age/:sex", _Uri) //需要加上三个动态参数，这三个参数将会返回到userInfo结构体中

	r.POST("/form", _Form)

	r.Run(":80")
}
