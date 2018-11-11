package route

import (
	"logging/api"
	"logging/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRoute(logging *logrus.Logger) *gin.Engine {
	route := gin.Default()
	route.Use(logger.LoggingMiddleware(logging, logger.NewUUID))
	route.GET("api/v1/user", api.ListUserHandler)
	route.GET("api/v1/user/:id", api.GetUserHandler)
	route.POST("api/v1/user", api.CreateUserHandler)
	return route
}
