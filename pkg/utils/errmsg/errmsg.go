package errmsg

import "github.com/gin-gonic/gin"

const (
	SUCCESS             = 0
	ERROR               = -1
	ERROR_INVALID_PARAM = 2

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code= 2000... 文章模块的错误

	ERROR_ARTICLE_NOT_EXIST = 2001
	ERROR_ARTICLE_SLUG_EXIST = 2002
	// code= 3000... 分类模块的错误
	ERROR_CATENAME_SLUG_EXIST = 3001
	ERROR_CATE_NOT_EXIST      = 3002
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "服务端发生未知错误",
	ERROR_INVALID_PARAM:    "参数错误",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
	ERROR_ARTICLE_SLUG_EXIST: "文章别名已存在",

	ERROR_CATENAME_SLUG_EXIST: "分类别名已存在",
	ERROR_CATE_NOT_EXIST:      "该分类不存在",
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
