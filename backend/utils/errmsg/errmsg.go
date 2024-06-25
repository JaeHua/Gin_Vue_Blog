package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//  code=1000...用户模块的错误

	ERROR_USERNAME_USED   = 1001
	ERROR_PASSWORD_WORNG  = 1002
	ERROR_USER_NOT_EXIST  = 1003
	ERROR_TOKEN_NOT_EXIST = 1004
	ERROR_TOKEN_RUNTIME   = 1005
	ERROR_TOKEN_WRONG     = 1006
	//code=2000...文章模拟的错误

	//code=3000...用户模块的错误
)

var codemsg = map[int]string{

	SUCCESS:               "OK",
	ERROR:                 "FAIL",
	ERROR_USERNAME_USED:   "用户已存在",
	ERROR_PASSWORD_WORNG:  "密码错误",
	ERROR_USER_NOT_EXIST:  "用户不存在",
	ERROR_TOKEN_NOT_EXIST: "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:   "TOKEN已过期",
	ERROR_TOKEN_WRONG:     "TOKEN错误",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
