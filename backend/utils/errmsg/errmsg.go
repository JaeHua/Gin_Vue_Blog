package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//  code=1000...用户模块的错误

	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WORNG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008

	//  code=2000...文章模拟的错误
	ERROR_ART_NOT_EXIST = 2001

	//code=3000...分类模块的错误
	ERROR_CATEGORY_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002

	//code=4000...验证模块错误
	ERROR_VCODE_SEND    = 4001
	ERROR_VCODE_STORE   = 4002
	ERROR_VCODE_WRONG   = 4003
	ERROR_VCODE_EXPIRED = 4004
)

var codemsg = map[int]string{

	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户已存在",
	ERROR_PASSWORD_WORNG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN错误",
	ERROR_USER_NO_RIGHT:    "该用户无权限",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_CATEGORY_USED:    "该分类已存在",
	ERROR_ART_NOT_EXIST:    "文章不存在",
	ERROR_CATE_NOT_EXIST:   "分类不存在",
	ERROR_VCODE_SEND:       "验证码发送失败",
	ERROR_VCODE_STORE:      "验证码存储失败",
	ERROR_VCODE_WRONG:      "验证码错误",
	ERROR_VCODE_EXPIRED:    "验证码失效",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
