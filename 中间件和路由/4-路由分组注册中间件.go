package main

import "github.com/gin-gonic/gin"

type Users struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Res struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

// Middleware 是 userControl 的一个全局中间件，需要在请求头上面加上一个 token 的值为 1234，否则就校验失败不会响应数据
func Middleware(c *gin.Context) {
	token := c.GetHeader("token")

	if token == "1234" {
		c.Next()
		return
	}

	c.JSON(200, Res{0, nil, "权限校验失败"}) //走到这一步就表示权限校验失败了，就不能接续往下走了
	c.Abort()                          //没有权限，拦截下来
}

//func Middleware(msg string) gin.HandlerFunc{ //上面的另一种写法，这样就可以传递参数，灵活一些
/*
	这里的代码一执行程序就会运行
*/

//这里的代码等请求来了才会运行
//	return func(c *gin.Context) {
//		token := c.GetHeader("token")
//
//		if token == "1234" {
//			c.Next()
//			return
//		}
//
//		c.JSON(200, Res{0, nil, msg})
//		c.Abort()
//	}
//}

func _userListView(c *gin.Context) {
	userlist := []Users{
		{"cfd", 23},
		{"xzk", 20},
		{"cdy", 21},
	}

	c.JSON(200, Res{0, userlist, "成功"})
}

func _UserRouterInit(router *gin.RouterGroup) {
	userControl := router.Group("user_manger")
	userControl.Use(Middleware) //这个一个分组的全局中间件
	{
		userControl.GET("/users", _userListView) //路由为 /api/user_manger/users
	}
}

func main() {
	r := gin.Default()

	api := r.Group("api") //路由分组，这一个括号里面的请求就表示是这一个分组的，需要加上/api。路由分组是可以嵌套的

	_UserRouterInit(api) //也可以封装成一个函数，放在其他包，这样就很清晰和简洁

	r.Run(":80")
}
