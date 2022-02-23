package errmsg

import "github.com/gin-gonic/gin"

const (
	SUCCESS           = 0
	ERROR             = -1
	ErrorNotInstalled = 2
	ErrorInvalidParam = 3

	// code= 1000... 用户模块的错误
	ErrorUsernameExist  = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenExist     = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
	ErrorUserNoRight    = 1008

	// code= 2000... 文章模块的错误
	ErrorArticleNotExist  = 2001
	ErrorArticleSlugExist = 2002

	// code= 3000... 分类模块的错误
	ErrorCategorySlugExist = 3001
	ErrorCateNotExist      = 3002
)

var codeMsg = map[int]string{
	SUCCESS:             "OK",
	ERROR:               "服务端发生未知错误",
	ErrorInvalidParam:   "参数错误",
	ErrorNotInstalled:   "未安装,跳转到安装页面",
	ErrorUsernameExist:  "用户名已存在！",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenExist:     "TOKEN不存在,请重新登陆",
	ErrorTokenRuntime:   "TOKEN已过期,请重新登陆",
	ErrorTokenWrong:     "TOKEN不正确,请重新登陆",
	ErrorTokenTypeWrong: "TOKEN格式错误,请重新登陆",
	ErrorUserNoRight:    "该用户无权限",

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
		"status": Success,
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
		"status": Error,
		"msg":    GetErrMsg(ERROR),
	}
}
