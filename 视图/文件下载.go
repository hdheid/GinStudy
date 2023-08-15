package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/download", func(c *gin.Context) {
		c.File("upload/哆啦A梦.jpg") //直接响应一个路径下的文件
	})

	//前后端模式下的文件下载

	r.Run(":80")
}
