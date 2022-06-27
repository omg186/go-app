package error_code

const (
	SUCCESS            = 0
	FAILURE            = 1
	StatusUnauthorized = 401
	NotFound           = 404
	ParamBindError     = 10000
	ParamRequired      = 10001
	ParamInvalid       = 10002
	ServerError        = 10101
	TooManyRequests    = 10102
	AuthorizationError = 10103
	RBACError          = 10104
	LoginIsInvalid     = 10105
)

func Text(code int) (str string) {
	lang := "zh_CN" //config.Config.Language

	var ok bool
	switch lang {
	case "zh_CN":
		str, ok = zhCNText[code]
		break
	case "en":
		str, ok = enUSText[code]
		break
	}
	if !ok {
		return "unknown error"
	}
	return
}
