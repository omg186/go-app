package response

import (
	"net/http"
	"order-food-service/pkg/error_code"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Resp() *Response {
	// 初始化response
	return &Response{
		httpCode: http.StatusOK,
		result: &Result{
			Code:    "0",
			Message: "",
			Data:    nil,
			Cost:    "",
		},
	}
}

// Success 业务成功响应
func Success(c *gin.Context, data ...any) {
	if data != nil {
		Resp().WithDataSuccess(c, data[0])
		return
	}
	Resp().Success(c)
}

// Fail 业务失败响应
func Fail(c *gin.Context, code string, message string, data ...any) {
	if data != nil {
		Resp().WithData(data[0]).FailCode(c, code, message)
		return
	}
	Resp().FailCode(c, code, message)
}

type Result struct {
	Code    string      `json:"code"`    // 业务code 0 代表成功
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 成功返回的数据
	Cost    string      `json:"cost"`
}

type Response struct {
	httpCode int
	result   *Result
}

// Fail 错误返回
func (r *Response) Fail(c *gin.Context, msg ...string) {
	r.SetCode(error_code.FAILURE)
	if msg != nil {
		r.WithMessage(msg[0])
	}
	r.json(c)
}

// FailCode 自定义错误码返回
func (r *Response) FailCode(c *gin.Context, code string, msg ...string) {
	// var newCode string
	codeInt, err := strconv.Atoi(code)
	if err == nil {
		if codeInt <= 500 {
			r.SetHttpCode(codeInt)
		}
	}
	r.SetCode(code)
	if msg != nil {
		r.WithMessage(msg[0])
	}
	r.json(c)
}

// Success 正确返回
func (r *Response) Success(c *gin.Context) {
	r.SetCode(error_code.SUCCESS)
	r.json(c)
}

// WithDataSuccess 成功后需要返回值
func (r *Response) WithDataSuccess(c *gin.Context, data interface{}) {
	r.SetCode(error_code.SUCCESS)
	r.WithData(data)
	r.json(c)
}

// SetCode 设置返回code码
func (r *Response) SetCode(code string) *Response {
	r.result.Code = code
	return r
}

// SetHttpCode 设置http状态码
func (r *Response) SetHttpCode(code int) *Response {
	r.httpCode = code
	return r
}

// type defaultRes struct {
// 	Result any `json:"result"`
// }

// WithData 设置返回data数据
func (r *Response) WithData(data interface{}) *Response {
	switch data.(type) {
	case string, int, bool:
		r.result.Data = data
	default:
		r.result.Data = data
	}
	return r
}

// WithMessage 设置返回自定义错误消息
func (r *Response) WithMessage(message string) *Response {
	r.result.Message = message
	return r
}

// json 返回 gin 框架的 HandlerFunc
func (r *Response) json(c *gin.Context) {
	if r.result.Message == "" {
		r.result.Message = error_code.Text(r.result.Code)
	}

	// if r.Data == nil {
	// 	r.Data = struct{}{}
	// }

	r.result.Cost = time.Since(c.GetTime("requestStartTime")).String()
	c.AbortWithStatusJSON(r.httpCode, r.result)
}
