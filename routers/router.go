package routers

import (
	"app/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouters() *gin.Engine {
	var r *gin.Engine
	if config.Config.Debug == false {
		// 生产模式
		//r = ReleaseRouter()
		//r.Use(
		//	middleware.RequestCostHandler(),
		//	middleware.CustomLogger(),
		//	middleware.CustomRecovery(),
		//	middleware.CorsHandler(),
		//)
	} else {
		// 开发调试模式
		r = gin.New()
		//r.Use(
		//	middleware.RequestCostHandler(),
		//	gin.Logger(),
		//	middleware.CustomRecovery(),
		//	middleware.CorsHandler(),
		//)
	}
	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	// 设置 API 路由
	setApiRoute(r)
	return r
}
