package routers

import (
	"order-food-service/api"

	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.RouterGroup) {
	r.POST("/login", api.Login)
}
