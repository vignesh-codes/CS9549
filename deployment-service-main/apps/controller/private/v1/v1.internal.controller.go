package v1

import (
	v1Internal "deployment-service/apps/dao/private/v1"
	"deployment-service/apps/repository/adapter"
	model_event_logger "deployment-service/models/model.eventlogger"
	"deployment-service/utils"

	"github.com/gin-gonic/gin"
)

type EventLoggerController struct {
	v1EventLoggerDao v1Internal.IEventLoggerDao
}

type IEventLoggerController interface {
	LogActivity(ctx *gin.Context)
}

func NewEventLoggerController(repository *adapter.Repository) IEventLoggerController {
	return &EventLoggerController{
		v1EventLoggerDao: v1Internal.NewEventLoggerDao(repository),
	}
}

func (ctrl EventLoggerController) LogActivity(ctx *gin.Context) {
	var request *model_event_logger.Event
	if ok := utils.BindJSON(ctx, &request); !ok {
		ctx.Abort()
		return
	}
	// event_source, found := ctx.Get("event_source")
	// if !found {
	// 	status := response.UnAuthorized(string(response.ErrUnauthorized))
	// 	ctx.JSON(status.Status(), status)
	// 	ctx.Abort()
	// 	return
	// }
	// request.EventSource = event_source.(string)
	request.Id = int(utils.GenerateUUID(request.EventSource))
	ctrl.v1EventLoggerDao.LogActivity(ctx, *request)
}
