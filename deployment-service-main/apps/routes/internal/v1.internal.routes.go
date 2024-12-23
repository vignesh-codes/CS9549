package internal

import (
	v1 "deployment-service/apps/controller/private/v1"
	"deployment-service/apps/repository/adapter"
	"deployment-service/middlewares"

	"deployment-service/logger"

	"github.com/gin-gonic/gin"
)

func V1(group *gin.RouterGroup, repository *adapter.Repository) {
	logger.ConsoleLogger.Debug("Initialising v1 internal group routes.")
	v1PrivateEventLoggerCtrl := v1.NewEventLoggerController(repository)
	group.Use(middlewares.ValidateHeaderSecrets(repository))
	{
		group.POST("/log/", v1PrivateEventLoggerCtrl.LogActivity)
	}
}
