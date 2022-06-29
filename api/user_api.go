package api

import (
	"order-food-service/pkg/error_code"
	"order-food-service/pkg/response"
	"order-food-service/pkg/utils"
	"order-food-service/service"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// @Tags User
// @Name getUserInfo
// @Summary 查询登录用户信息
// @Param Authorization header string true "Authorization"
// @Success 200 {object} response.Result{data=dto.UserRes} "成功返回"
// @Accept json
// @Produce json
// @Router /userInfo [GET]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	tokenInfo, _ := c.Get("claims")
	errText := error_code.Text(error_code.StatusUnauthorized)
	if tokenInfo == nil {
		response.Resp().SetCode(error_code.StatusUnauthorized).Fail(c, errText)
	}
	userInfo, err := tokenInfo.(*utils.CustomClaims)
	if !err {
		response.Resp().SetCode(error_code.StatusUnauthorized).Fail(c, errText)
	}
	userService := service.UserService{}
	userRes, userErr := userService.GetUserNamePassword(userInfo.UserName, userInfo.Password)
	if userErr != nil {
		response.Resp().SetCode(userErr.Error()).Fail(c, errText)
	}
	response.Success(c, userRes)
}
