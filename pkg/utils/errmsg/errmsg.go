package errmsg

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

	// code= 4000... 评论模块的错误

	// code= 5000... 附件模块的错误
	ErrorOnlySupportZip = 5001
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

	ErrorOnlySupportZip: "仅支持zip格式",
}

type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}

func Success(data interface{}) Response {
	return Response{
		Status: SUCCESS,
		Msg:    GetErrMsg(SUCCESS),
		Data:   data,
	}
}

func Fail(status int) Response {
	return Response{
		Status: status,
		Msg:    GetErrMsg(status),
	}
}

func FailMsg(msg string) Response {
	return Response{
		Status: ERROR,
		Msg:    msg,
	}
}

func Error() Response {
	return Response{
		Status: ERROR,
		Msg:    GetErrMsg(ERROR),
	}
}
