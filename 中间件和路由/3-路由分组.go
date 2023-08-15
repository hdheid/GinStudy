package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func userListView(c *gin.Context) {
	userlist := []UserInfo{
		{"cfd", 23},
		{"xzk", 20},
		{"cdy", 21},
	}

	c.JSON(200, Response{0, userlist, "成功"})
}

func UserRouterInit(router *gin.RouterGroup) {
	userControl := router.Group("user_manger")
	{
		userControl.GET("/users", userListView) //路由为 /api/user_manger/users
	}
}

func main() {
	r := gin.Default()

	api := r.Group("api") //路由分组，这一个括号里面的请求就表示是这一个分组的，需要加上/api。路由分组是可以嵌套的
	{
		userControl := api.Group("user_manger")
		{
			userControl.GET("/users", userListView) //路由为 /api/user_manger/users
		}

		// UserRouterInit(api) //也可以封装成一个函数，放在其他包，这样就很清晰和简洁
	}

	r.Run(":80")
}
