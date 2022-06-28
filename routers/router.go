package routers

import (
	"net/http"
	"order-food-service/middleware"
	"order-food-service/pkg/error_code"
	"order-food-service/pkg/response"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.New()
	r.NoRoute(func(context *gin.Context) {
		response.Fail(context, http.StatusNotFound, error_code.Text(http.StatusNotFound))
	})
	r.NoMethod(func(context *gin.Context) {
		response.Fail(context, http.StatusMethodNotAllowed, error_code.Text(http.StatusNotFound))
	})
	unauthorizedGroup := r.Group("/api")
	{
		unauthorizedGroup.Use(cors.Default())
		unauthorizedGroup.Use(middleware.LoggerToFile())

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
