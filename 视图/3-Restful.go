package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"reflect"
	"strconv"
)

type ArticleModel struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Response 将数据封装一下
type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

// DataInit 设置全局变量
func DataInit() []ArticleModel {
	articleList := []ArticleModel{
		{1, "语文", "语文书"},
		{2, "数学", "数学书"},
		{3, "英语", "英语书"},
	}
	return articleList
}

func GetList(c *gin.Context) {
	articleList := DataInit()

	c.JSON(200, Response{
		Code:    0,
		Data:    articleList,
		Message: "查找成功",
	})
}

func GetDetail(c *gin.Context) {
	_, book := Find(c) //封装成了函数，通过获取动态ID查找相应的数据并返回

	c.JSON(200, Response{
		Code:    0,
		Data:    book,
		Message: "查找成功",
	})
}

// Create 接收前端传来的json数据
func Create(c *gin.Context) {
	_, book := BindJson(c) //封装成了函数，将从前端获取到的JSON数据存储到结构体中

	fmt.Println(book) //实际是将获取到的数据传到数据库中，在这里输出展示

	c.JSON(200, Response{ //将接收成功的数据传回
		Code:    0,
		Data:    book,
		Message: "添加成功",
	})
}

func Update(c *gin.Context) {
	articleList := DataInit()
	id, _ := strconv.Atoi(c.Param("id"))
	_, NewBook := BindJson(c)

	for i, _ := range articleList { //找到需要修改的数据，进行修改
		if articleList[i].Id == id {
			articleList[i].Title = NewBook.Title
			articleList[i].Content = NewBook.Content
		}
	}

	c.JSON(200, Response{ //将接收成功的数据传回
		Code:    0,
		Data:    articleList,
		Message: "修改成功",
	})
}

func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	articleList := DataInit()

	newArticleList := make([]ArticleModel, 0)
	for _, book := range articleList {
		if book.Id != id {
			newArticleList = append(newArticleList, book)
		}
	}
	articleList = newArticleList //将等于 id 的删除

	c.JSON(200, Response{
		Code:    0,
		Data:    articleList,
		Message: "删除成功",
	})
}

func (book ArticleModel) IsEmpty() bool { //结构体判空函数
	return reflect.DeepEqual(book, ArticleModel{})
}

// Find 通过ID查询数据的函数
func Find(c *gin.Context) (error, ArticleModel) {
	//获取param中的id,正常情况下是进行数据库的查询，在这里简化一下
	id, _ := strconv.Atoi(c.Param("id"))
	articleList := DataInit()

	var bk ArticleModel
	for _, book := range articleList {
		if book.Id == id {
			bk = book
			break
		}
	}
	if bk.IsEmpty() {
		return errors.New("ID NOT FOUND"), bk
	}

	return nil, bk
}

// BindJson 将前端JSON数据提取出来保存到结构体中
func BindJson(c *gin.Context) (error, ArticleModel) {
	body, _ := c.GetRawData()
	ContentType := c.GetHeader("Content-Type")

	var book ArticleModel
	switch ContentType {
	case "application/json":

		err := json.Unmarshal(body, &book)
		if err != nil {
			return errors.New("BIND FAILED"), book
		}
	}

	return nil, book
}

func main() {
	r := gin.Default()

	r.GET("/articles", GetList)       //获取文章列表
	r.GET("/articles/:id", GetDetail) //获取文章详情
	r.POST("/articles", Create)       //添加文章
	r.PUT("/articles/:id", Update)    //编辑文章
	r.DELETE("/articles/:id", Delete) //删除文章

	r.Run(":80")
}
