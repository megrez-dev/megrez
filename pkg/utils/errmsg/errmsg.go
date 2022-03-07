package errmsg

import "github.com/gin-gonic/gin"

const (
	SUCCESS           = 0
	ERROR             = -1
	ErrorNotInstalled = 2
	ErrorInvalidParam = 3

	// code= 1000... 用户模块的错误
	ErrorUsernameExist = 1001
	ErrorPasswordWrong = 1002
	ErrorUserNotExist  = 1003
	ErrorTokenNotExist = 1004
	ErrorTokenInvalid  = 1005
	ErrorTokenExpired  = 1006

	// code= 2000... 文章模块的错误
	ErrorArticleNotExist  = 2001
	ErrorArticleSlugExist = 2002

	// code= 3000... 分类模块的错误
	ErrorCategorySlugExist = 3001
	ErrorCateNotExist      = 3002
)

var codeMsg = map[int]string{
	SUCCESS:            "OK",
	ERROR:              "服务端发生未知错误",
	ErrorInvalidParam:  "请求参数错误",
	ErrorNotInstalled:  "博客未安装,请先完成安装程序",
	ErrorUsernameExist: "用户名已存在！",
	ErrorPasswordWrong: "用户名或密码错误",
	ErrorUserNotExist:  "用户不存在",
	ErrorTokenNotExist: "请登录后再操作",
	ErrorTokenInvalid:  "认证失败，请重新登录",
	ErrorTokenExpired:  "认证过期，请重新登录",

	ErrorArticleNotExist:  "文章不存在",
	ErrorArticleSlugExist: "文章别名已存在",

	ErrorCategorySlugExist: "分类别名已存在",
	ErrorCateNotExist:      "该分类不存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

func Success(data interface{}) gin.H {
	return gin.H{
		"status": SUCCESS,
		"msg":    GetErrMsg(SUCCESS),
		"data":   data,
	}
}

func Fail(status int) gin.H {
	return gin.H{
		"status": status,
		"msg":    GetErrMsg(status),
	}
}

func Error() gin.H {
	return gin.H{
		"status": ERROR,
		"msg":    GetErrMsg(ERROR),
	}
}
