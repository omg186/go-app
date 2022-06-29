package error_code

const (
	SUCCESS                = "0"
	FAILURE                = "1"
	StatusUnauthorized     = "401"
	ParamRequired          = "400"
	NotFound               = "404"
	StatusMethodNotAllowed = "405"
	ServerError            = "500"
	ParamBindError         = "10000"
	TooManyRequests        = "10102"
	AuthorizationError     = "10103"
	RBACError              = "10104"
	LoginIsInvalid         = "10105"
	UserNotFound           = "20000"
)

func Text(code string) (str string) {
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
		return code + "unknown error"
	}
	return
}
