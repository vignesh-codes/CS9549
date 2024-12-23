package v1

import (
	"deployment-service/apps/repository/adapter"
	"deployment-service/apps/svc"
	model_deployment "deployment-service/models/model.deployment"
	"fmt"
	"net/http"

	"gorm.io/gorm/utils"

	"github.com/gin-gonic/gin"
)

type DeploymentDao struct {
	ServiceRepo *svc.ServiceRepository
}

type IDeploymentsDao interface {
	GetDeployments(ctx *gin.Context, namespace string)
	GetTenantKubernetesInfo(ctx *gin.Context, namespace string)
	GetDeploymentByName(ctx *gin.Context, namespace string)
	CreateNamespace(ctx *gin.Context, namespace string)
	CreateDeployment(ctx *gin.Context, payload *model_deployment.CreateDeploymentRequest)
	DeleteDeployment(ctx *gin.Context, namespace string, deploymentName string)
	GetLatestEvents(ctx *gin.Context, namespace string, topK int)
	UpdateDeploymentByName(ctx *gin.Context, namespace string, payload *model_deployment.UpdateDeploymentReq)
}

func NewDeploymentsDao(repository *adapter.Repository) IDeploymentsDao {
	return &DeploymentDao{
		ServiceRepo: svc.NewServiceRepo(repository),
	}
}

func (dao DeploymentDao) GetLatestEvents(ctx *gin.Context, namespace string, topK int) {
	events, err := dao.ServiceRepo.DeploymentService.GetLatestEvents(namespace, topK)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, events)
	ctx.Abort()
}

func (dao DeploymentDao) GetDeployments(ctx *gin.Context, namespace string) {
	var resp = map[string]interface{}{}
	db_response, err := dao.ServiceRepo.DeploymentService.GetAllDeploymentsFromDBByNamespace(namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		ctx.Abort()
		return
	}
	for i, db_resp := range db_response {
		kubernetes_resp, err := dao.ServiceRepo.DeploymentService.GetDeploymentByName(namespace, db_resp.Name)
		if err != nil {
			fmt.Println("error getting kubernetes resp 40 ", err)
		}
		resp[utils.ToString(i)] = map[string]interface{}{
			"deployment_info": db_resp,
			"kubernetes_resp": kubernetes_resp,
		}
	}
	ctx.JSON(http.StatusOK, resp)
	ctx.Abort()
}

func (dao DeploymentDao) GetTenantKubernetesInfo(ctx *gin.Context, namespace string) {

	response, err := dao.ServiceRepo.DeploymentService.GetTenantKubernetesInfo(namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
}

func (dao DeploymentDao) GetDeploymentByName(ctx *gin.Context, namespace string) {
	response, err := dao.ServiceRepo.DeploymentService.GetDeploymentByName(namespace, ctx.Param("deployment_name"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"message": response})
	ctx.Abort()
}

func (dao DeploymentDao) CreateNamespace(ctx *gin.Context, namespace string) {
	response := dao.ServiceRepo.DeploymentService.CreateNamespaceIfNotExists(namespace)

	ctx.JSON(http.StatusOK, map[string]interface{}{"message": "successfully created namespace", "err": response})
	ctx.Abort()
}

func (dao DeploymentDao) CreateDeployment(ctx *gin.Context, payload *model_deployment.CreateDeploymentRequest) {
	resp, err := dao.ServiceRepo.DeploymentService.CreateDeployment(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Successfully Created Deployment: %s", payload.Name),
		"result":  resp})
	ctx.Abort()
}

func (dao DeploymentDao) DeleteDeployment(ctx *gin.Context, namespace, deployment_name string) {
	deployment_info, err := dao.ServiceRepo.DeploymentService.GetDeploymentFromDBByName(namespace, deployment_name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		ctx.Abort()
		return
	}
	_, err = dao.ServiceRepo.DeploymentService.DeleteDeployment(namespace, deployment_name, deployment_info.RepoScoutId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"message": fmt.Sprintf("Successfully Deleted Deployment: %s", deployment_name)})
	ctx.Abort()
}

func (dao DeploymentDao) UpdateDeploymentByName(ctx *gin.Context, namespace string, payload *model_deployment.UpdateDeploymentReq) {
	fmt.Println("updateing deployment ")
	resp, err := dao.ServiceRepo.DeploymentService.UpdateDeploymentByName(namespace, payload.Name, payload.Image, payload.Replicas)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"message": err.Error()})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Successfully Updated Deployment: %s", resp),
		"result":  resp})
	ctx.Abort()
}
