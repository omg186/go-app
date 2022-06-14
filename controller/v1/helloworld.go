package v1

import (
	"app/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func HelloWorld(c *gin.Context) {
	str, ok := c.GetQuery("name")
	if !ok {
		str = "gin1"
	}
	response.Success(c, fmt.Sprintf("hello %s", str))
}
