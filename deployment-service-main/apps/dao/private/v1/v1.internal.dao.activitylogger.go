package v1

import (
	"deployment-service/apps/repository/adapter"
	"deployment-service/apps/svc"
	"net/http"

	model_event_logger "deployment-service/models/model.eventlogger"

	"github.com/gin-gonic/gin"
)

type EventLoggerDao struct {
	ServiceRepo *svc.ServiceRepository
}

type IEventLoggerDao interface {
	LogActivity(ctx *gin.Context, payload model_event_logger.Event)
}

func NewEventLoggerDao(repository *adapter.Repository) IEventLoggerDao {
	return &EventLoggerDao{
		ServiceRepo: svc.NewServiceRepo(repository),
	}
}

func (dao EventLoggerDao) LogActivity(ctx *gin.Context, payload model_event_logger.Event) {
	var response []map[string]interface{}

	response, _ = dao.ServiceRepo.EventLoggerService.LogEvent(payload)

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": response})
	ctx.Abort()
}
