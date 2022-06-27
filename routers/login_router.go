package routers

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

func LoginRouter(r *gin.RouterGroup) {
	r.POST("/login", api.Login)
}
