package api

import (
	"order-food-service/dto"
	"order-food-service/global"
	"order-food-service/pkg/error_code"
	"order-food-service/pkg/response"
	"order-food-service/pkg/utils"
	"order-food-service/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginApi struct {
}
type LoginRes struct {
	Token string `json:"token"` // token
}

// @Tags Login
// @Name LoginUser
// @Summary 用户名密码登陆
// @Param body body dto.LoginRequest true "登陆参数"
// @Success 200 {object} response.Result{data=LoginRes} "成功返回"
// @Failure 400 {object} response.Result{message=string} "参数错误"
// @Accept json
// @Produce json
// @Router /login [post]
func (l *LoginApi) Login(c *gin.Context) {
	var json dto.LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		global.AppLog.Warning(err)
		// c.JSON(400, nil)
		code, err := utils.GetError(err.(validator.ValidationErrors))
		response.Resp().FailCode(c, code, err)
		// response.Fail(c, 400, "Login failed")
		return
	}
	loginService := service.LoginService{}
	token, err := loginService.Login(json.UserName, json.PassWord)
	if err != nil {
		response.Resp().FailCode(c, err.Error(), error_code.Text(err.Error()))
		return
	}
	response.Success(c, LoginRes{Token: token})
}

// @Summary 退出登录
// @Description 退出登录
// @Tags Posts
// @Produce json
// @Success 200 {string} string "ok" "返回用户信息"
// @Router /api/OutLogin [post]
func OutLogin(c *gin.Context) {

}
