package routers

import (
	"order-food-service/middleware"
	"order-food-service/pkg/error_code"
	"order-food-service/pkg/response"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	// r := gin.New()
	r := gin.Default()
	// r.
	r.NoRoute(func(context *gin.Context) {
		response.Fail(context, error_code.NotFound, error_code.Text(error_code.NotFound))
	})
	r.NoMethod(func(context *gin.Context) {
		response.Fail(context, error_code.StatusMethodNotAllowed, error_code.Text(error_code.StatusMethodNotAllowed))
	})
	unauthorizedGroup := r.Group("/api")
	{
		unauthorizedGroup.Use(cors.Default())
		unauthorizedGroup.Use(middleware.LoggerToFile())

		SwaggerRouter(unauthorizedGroup)
		PingRouter(unauthorizedGroup)
		LoginRouter(unauthorizedGroup)
	}
	// LoginPwdRouter(unauthorizedGroup)
	// SwaggerRouter(unauthorizedGroup)

	// authorizedGroup := r.Group("/api", gin.BasicAuth(gin.Accounts{
	// 	"foo":    "bar",
	// 	"austin": "1234",
	// 	"lena":   "hello2",
	// 	"manu":   "4321",
	// }))
	authorizedGroup := r.Group("/api")
	{
		authorizedGroup.Use(cors.Default())
		authorizedGroup.Use(middleware.LoggerToFile())
		authorizedGroup.Use(middleware.JWTAuth())

		PingAuthRouter(authorizedGroup)
	}
	// SysDictTypeRouter(authorizedGroup)
	// SysDictDataRouter(authorizedGroup)
	// DistrictRouter(authorizedGroup)
	return r
}
