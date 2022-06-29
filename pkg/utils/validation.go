package utils

import (
	"fmt"
	"order-food-service/pkg/error_code"

	"github.com/go-playground/validator/v10"
)

// 绑定模型获取验证错误的方法
func GetError(err error) (string, string) {
	fieldError, ok := err.(validator.ValidationErrors)
	if !ok {
		return error_code.ParamBindError, error_code.Text(error_code.ParamBindError)
		// return ParamBindError, fmt.Sprintf(global.ERR_MSG[global.ERROR_CODE_PARAM_INVALID], "")
	}
	for _, v := range fieldError {
		switch v.ActualTag() {
		case "required":
			return error_code.ParamRequired, fmt.Sprintf(error_code.Text(error_code.ParamRequired), v.Field())
		case "gte", "gt", "ne", "lt", "lte":
			return error_code.ParamRequired, fmt.Sprintf(error_code.Text(error_code.ParamRequired), v.Field())

			// return global.ERROR_CODE_PARAM_INVALID, fmt.Sprintf(global.ERR_MSG[global.ERROR_CODE_PARAM_INVALID], v.Field())
		}
	}
	return error_code.ParamRequired, fmt.Sprintf(error_code.Text(error_code.ParamRequired), "")
}
