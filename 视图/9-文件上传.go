package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func CreateCopy(c *gin.Context) {
	file, _ := c.FormFile("file")
	readerFile, _ := file.Open()                 //将上传文件打开
	writerFile, _ := os.Create("./upload/2.jpg") //创建一个新文件
	defer writerFile.Close()
	n, _ := io.Copy(writerFile, readerFile) //将 readerFile 中的内容复制到 writerFile 中去
	fmt.Println(n)
	c.JSON(200, gin.H{"msg": "成功"})
}

func main() {
	r := gin.Default()

	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file") //上传的文件名，基本使用form-data上传文件
		fmt.Println(file.Filename)    //文件名
		fmt.Println(file.Size)        //文件大小，单位是字节

		//服务端保存文件的方式之一
		c.SaveUploadedFile(file, "upload/1.jpg") //第二个表示上传路径，从根目录开始，上传到xxx.jpg，这样照片就会以xxx.jpg的形式保留下来

		c.JSON(200, gin.H{"msg": "上传成功"})
	})

	r.POST("/upload2", CreateCopy)

	r.Run(":80")
}
