package routers

import (
	"order-food-service/api"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	apiUser := api.UserApi{}
	r.GET("/getUserInfo", apiUser.GetUserInfo)
}
