package api

import (
	"github.com/gin-gonic/gin"
)

// @Summary 测试SayHello
// @Description 向你说Hello
// @Id Ping
// @Tags Ping
// @Accept mpfd
// @Produce json
// @Param who query string true "人名"
// @Router /ping [get]
// x-{ "tags": "getDog", "operationId": "getDog" }
func Ping(c *gin.Context) {
	// c.da
	c.JSON(200, gin.H{"message": "pong"})
}
