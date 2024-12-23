package svc

import (
	adapter "deployment-service/apps/repository/adapter"
)

var ServiceRepo *ServiceRepository

type ServiceRepository struct {
	EventLoggerService *EventLoggerService
	DeploymentService  *DeploymentService
	BuildService       *BuildService
}

func NewServiceRepo(repository *adapter.Repository) *ServiceRepository {
	return &ServiceRepository{
		EventLoggerService: &EventLoggerService{repository},
		DeploymentService:  &DeploymentService{repository},
		BuildService:       &BuildService{repository},
	}
}
