package routers

import (
	c "app/controller/v1"
	"github.com/gin-gonic/gin"
)

func setApiRoute(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/hello-world", c.HelloWorld)
	}
}
