package api

import (
	"net/http"
	"order-food-service/pkg/response"
	"order-food-service/pkg/utils"
	"order-food-service/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// @Tags User
// @Name GetUserInfo
// @Summary 用户名密码登陆
// @Param Authorization header string true "Authorization" "Token"
// @Success 200 {object} response.Result{data=dto.UserRes} "成功返回"
// @Failure 401 {object} response.Result "token无效"
// @Accept json
// @Produce json
// @Router /login [post]
func (u *UserApi) GetUserInfo(c *gin.Context) {
	tokenInfo, _ := c.Get("claims")
	if tokenInfo != nil {
		response.Resp().SetCode(strconv.Itoa(http.StatusUnauthorized)).Fail(c, "token 失效")
	}
	userInfo, err := tokenInfo.(*utils.CustomClaims)
	if !err {
		response.Resp().SetCode(strconv.Itoa(http.StatusUnauthorized)).Fail(c, "token 失效")
	}
	userService := service.UserService{}
	userRes, userErr := userService.GetUserNamePassword(userInfo.UserName, userInfo.Password)
	if userErr != nil {
		response.Resp().SetCode(strconv.Itoa(http.StatusUnauthorized)).Fail(c, "token 失效")
	}
	response.Success(c, userRes)
}
