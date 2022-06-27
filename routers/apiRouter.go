package routers

// import (
// 	c "app/controller/v1"
// 	"app/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func setApiRoute(r *gin.Engine) {
// 	// r.BasePath("api")

// 	api := r.Group("/api")
// 	{
// 		v1Auth := api.Group("/v1")
// 		{
// 			v1Auth.Use(middleware.JWTAuth())
// 			v1Auth.POST("/authping", c.HelloWorld)
// 			menu := v1Auth.Group("/menu")
// 			{
// 				menu.GET("/", c.MenuController{}.Get)
// 			}
// 		}
// 		v1NoAuth := api.Group("/v1")
// 		{
// 			v1NoAuth.POST("/ping", c.HelloWorld)
// 			v1NoAuth.POST("/login", c.LoginController{}.Login)
// 		}
// 	}
// }
