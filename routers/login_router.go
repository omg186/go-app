package routers

import (
	"order-food-service/api"

	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.RouterGroup) {
	apiLogin := api.LoginApi{}
	r.POST("/login", apiLogin.Login)
	r.GET("/loginout", apiLogin.LoginOut)
}
