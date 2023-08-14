package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

type Users struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名不合法"`
	Age  int    `json:"age" binding:"required" msg:"年龄不合法"`
}

// GetValidMsg 获取结构体中的 msg 参数
func _GetValidMsg(err error, user *Users) string {
	getObj := reflect.TypeOf(user)
	//将 err 接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			if f, _ok := getObj.Elem().FieldByName(e.Field()); _ok { //根据报错字段获取结构体的具体字段
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}

func signValid(fl validator.FieldLevel) bool {
	nameList := []string{"cfd", "xzk", "cdy", "xpz", "hxk"}

	for _, name := range nameList {
		if name == fl.Field().Interface().(string) {
			return false //只要人名在名字列表中出现过，就返回假表示校验失败
		}
	}

	return true
}

func main() {
	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	r.POST("/", func(c *gin.Context) {
		var user Users
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
	})

	r.Run(":80")
}
