package app

const (
	QUERY_STRING_ERROR = 1001
	DATABASE_ERROR     = 1003
	USER_NOT_EXIST     = 1004
	MISS_TOKEN			= 1005
	USER_LOGIN_ERROR 	= 1006
	DATA_NOT_FOUND = 1007
	AUTH_ERROR = 1008
	USER_NOT_LOGIN = 1009
)

var errMessage = map[int]string{
	QUERY_STRING_ERROR: "请求参数不完整或有误",
	DATABASE_ERROR:     "数据库异常",
	USER_NOT_EXIST:     "用户不存在",
	MISS_TOKEN:			"缺少token",
	USER_LOGIN_ERROR:	"登陆错误",
	DATA_NOT_FOUND:"数据不存在",
	AUTH_ERROR:"权限不足",
	USER_NOT_LOGIN:"用户未登陆",
}

func GetErrMessage(code int) string {
	v, ok := errMessage[code]
	if ok {
		return v
	}
	return ""
}
