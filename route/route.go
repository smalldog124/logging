package route

import (
	"logging/api"
	"logging/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func NewRoute(logging *logrus.Logger, apiUser api.UserAPI) *gin.Engine {
	route := gin.Default()
	route.Use(logger.LoggingMiddleware(logging, logger.NewUUID))
	route.GET("api/v1/user", apiUser.ListUserHandler)
	route.GET("api/v1/user/:id", apiUser.GetUserHandler)
	route.POST("api/v1/user", apiUser.CreateUserHandler)
	return route
}
