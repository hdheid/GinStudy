package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1...in")
	c.Next() //使用这个函数之后，会优先按照顺序执行后面的中间件，直到走完了最后一个中间件再走下面的。如果使用了Abort函数，就表示当前为最后一个中间件
	//c.Abort() //可以使用这个函数进行拦截，本来是会响应两次，这样就只会响应上面这一次
	fmt.Println("m2...out")
}

func m2(c *gin.Context) {
	fmt.Println("m1...in")
	c.Next()
	c.JSON(200, gin.H{"msg": "mi的响应"})
	fmt.Println("m2...out")
}

func m3(c *gin.Context) {
	fmt.Println("m1...in")
	c.Next()
	fmt.Println("m2...out")
}

func main() {
	r := gin.Default()

	r.GET("/", m1, m2, m3)
	//不止能够写一个func函数，能够写很多个func函数，这就是中间件

	r.Run(":80")
}
