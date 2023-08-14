package main

import "github.com/gin-gonic/gin"

//常用验证器,多个可以使用 ',' 隔开
//binging:"required" 意思是这一个数据为必填，如果没有填写或者为空，校验的时候就会校验失败

//针对字符串
//min=5 表示最小长度为5
//max=6 表示最大长度为6
//len=6 表示要求长度为6

//针对数字大小
//eq=3 表示要求等于3
//ne=12 表示要求不能等于12
//gt=10 表示要求大于10
//gte=10 表示要求大于等于10
//lt=10 表示要求小于10
//lte=10 表示要求小于等于10

//针对一个结构体里面的同级字段的，例如确认密码的判定
//eqfield 等于其他字段的值
//nefield 不等于其他字段的值

//- 忽略某个字段，例如：binding:"-"

// binding里面不能有空格，否则会造成500错误码返回
type SignUserInfo struct {
	Name       string   `json:"name" binding:"required"`                     //姓名
	Age        int      `json:"age"`                                         //年龄
	Password   string   `json:"password" binding:"required,min=6,max=12"`    //密码
	RePassword string   `json:"re-password" binding:"eqfield=Password"`      //确认密码
	Sex        string   `json:"sex" binding:"oneof=man woman"`               //性别
	LikeList   []string `json:"likeList" binding:"required,dive,contains=f"` //爱好列表
	IP         string   `json:"ip" binding:"ip"`
	Url        string   `json:"url" binding:"url"`
	Date       string   `json:"date" binding:"datetime=2006-01-02"`
}

//oneof 枚举，表示只能是几个列举的数据当中的一个，否则报错

//字符串操作
//contains = f 		表示字符串需要包含 f
//excludes 			与上面相反，不包含
//startswith = f 	字符串前缀,字符串的前缀为f
//endswith 			字符串后缀，同理

//数组操作
//dive,contains=f 表示数组的每一个元素都必须包含f。第二个验证器可以为其他的，写在dive后面的都必须让数组的每一个元素都验证

//网络验证
//ip 判断 IP 是否符合要求
//ipv4
//ipv6
//uri
//url 判断 url 是否符合规则

//日期验证
//datetime=2006-01-02 15:04:05 后面是不是具体时间，而是表示是这样一种格式。而且必须这样写，可以不写后面的时分秒表示只看年月日

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		var user SignUserInfo

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, gin.H{"msg": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": user})
	})

	r.Run(":80")
}
