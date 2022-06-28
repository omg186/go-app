package v1

import (
	"fmt"
	"order-food-service/pkg/response"

	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	str, ok := c.GetQuery("name")
	if !ok {
		str = "gin1"
	}
	response.Success(c, fmt.Sprintf("hello %s", str))
}
