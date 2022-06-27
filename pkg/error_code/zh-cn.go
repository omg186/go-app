package error_code

var zhCNText = map[int]string{
	SUCCESS:            "OK",
	FAILURE:            "FAIL",
	NotFound:           "资源不存在",
	LoginIsInvalid:     "用户名密码无效",
	ServerError:        "内部服务器错误",
	TooManyRequests:    "请求过多",
	ParamBindError:     "参数错误",
	ParamRequired:      "参数 '%s' 是必须的",
	ParamInvalid:       "参数 '%s' 无效",
	AuthorizationError: "权限错误",
	StatusUnauthorized: "token过期或没有token",
	RBACError:          "暂无访问权限",
}
