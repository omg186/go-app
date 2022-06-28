package api

import (
	"net/http"
	loginDto "order-food-service/dto/login_dto"
	"order-food-service/global"
	"order-food-service/pkg/response"
	"order-food-service/pkg/utils"
	"order-food-service/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// type LoginApi struct{}

// LoginApi
// @summary 登录
func Login(c *gin.Context) {
	var json loginDto.LoginRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		global.AppLog.Warning(err)
		// c.JSON(400, nil)
		code, err := utils.GetError(err.(validator.ValidationErrors))
		response.Resp().SetHttpCode(http.StatusBadRequest).FailCode(c, code, err)
		// response.Fail(c, 400, "Login failed")
		return
	}
	loginService := service.LoginService{}
	token, err := loginService.Login(json.UserName, json.PassWord)
	if err != nil {
		response.Resp().SetHttpCode(http.StatusInternalServerError).Fail(c, err.Error())
		return
	}
	response.Success(c, token)
}
