package routers

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

func PingRouter(r *gin.RouterGroup) {
	r.GET("/ping", api.Ping)
}

func PingAuthRouter(r *gin.RouterGroup) {
	r.GET("/ping-auth", api.Ping)
}
