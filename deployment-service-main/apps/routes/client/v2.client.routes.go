package client

import (
	v1 "deployment-service/apps/controller/client/v1"
	"deployment-service/apps/repository/adapter"
	"deployment-service/logger"
	"deployment-service/middlewares"
	"fmt"

	"github.com/gin-gonic/gin"
)

func V2(group *gin.RouterGroup, repository *adapter.Repository) {
	logger.ConsoleLogger.Debug("Initialising frontend v1 group routes.")
	fmt.Println("Initialising frontend v1 group routes.")
	v1ClientDeploymentsCtrl := v1.NewDeploymentController(repository)
	v1ClientBuildsCrtrl := v1.NewBuildController(repository)
	group.Use(middlewares.ValidateHeaderSecrets(repository))
	{
		group.POST("/deployments/createns/", middlewares.ValidateJWT(repository), v1ClientDeploymentsCtrl.CreateNamespace)
		group.GET("/deployments/", middlewares.ValidateJWT(repository), v1ClientDeploymentsCtrl.GetDeploymentsByNamespace)
		group.GET("/deployments/events/", middlewares.ValidateJWT(repository), v1ClientDeploymentsCtrl.GetLatestEvents)
		group.GET("/deployments/tenant/", middlewares.ValidateJWT(repository), v1ClientDeploymentsCtrl.GetTenantKubernetesInfo)
		// create a new deployment
		group.POST("/deployments/", middlewares.ValidateJWT(repository), v1ClientDeploymentsCtrl.CreateDeployment)
		// Update a deployment replica
		group.PUT("/deployments/", middlewares.ValidateJWT(repository), v1ClientDeploymentsCtrl.UpdateDeploymentByName)
		// get a deployment by name
		group.GET("/deployments/:deployment_name", middlewares.ValidateJWT(repository), v1ClientDeploymentsCtrl.GetDeploymentByName)
		// delete a deployment by name
		group.DELETE("/deployments/:deployment_name", middlewares.ValidateJWT(repository), v1ClientDeploymentsCtrl.DeleteDeployment)

		group.POST("/build/scout/", middlewares.ValidateJWT(repository), v1ClientBuildsCrtrl.CreateNewRepoScout)
		group.GET("/build/scout/", middlewares.ValidateJWT(repository), v1ClientBuildsCrtrl.GetAllRepoScouts)
	}
}
