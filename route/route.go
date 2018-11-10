package route

import (
	"logging/api"

	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	route := gin.Default()
	route.GET("api/v1/user", api.ListUser)
	return route
}
