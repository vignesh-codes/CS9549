package v1

import (
	"deployment-service/apps/repository/adapter"
	"deployment-service/apps/svc"
	"deployment-service/logger"
	model_build "deployment-service/models/model.build"
	"deployment-service/utils"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type BuildDao struct {
	ServiceRepo *svc.ServiceRepository
}

type IBuildDao interface {
	CreateNewRepoScout(ctx *gin.Context, request *model_build.RepoScout)
	GetAllRepoScouts(ctx *gin.Context, request *model_build.RepoScout)
}

func NewBuildDao(repository *adapter.Repository) IBuildDao {
	return &BuildDao{
		ServiceRepo: svc.NewServiceRepo(repository),
	}
}

func (dao BuildDao) CreateNewRepoScout(ctx *gin.Context, request *model_build.RepoScout) {
	response, _ := dao.ServiceRepo.BuildService.CreateNewRepoScout(*request)
	ctx.JSON(http.StatusOK, map[string]interface{}{"message": response})
	ctx.Abort()
}

func (dao BuildDao) GetAllRepoScouts(ctx *gin.Context, request *model_build.RepoScout) {
	// Get all repo scouts for the given namespace
	repoScouts, err := dao.ServiceRepo.BuildService.GetAllRepoScouts(request.Namespace)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error fetching repo scouts",
			"error":   err.Error(),
		})
		return
	}
	fmt.Println("got all repo scouts", repoScouts)
	// Prepare the response as a list of repo scouts
	response := []interface{}{}

	// Iterate over each repo scout
	for _, repoScout := range repoScouts {
		repoResponse := map[string]interface{}{
			"repo_name":        repoScout.RepoName,
			"repo_scout_id":    repoScout.ID,
			"latest_image_url": "",
			"deployments":      []map[string]interface{}{}, // Nested deployments data for each repo
			"release_info":     map[string]interface{}{},   // Release info data for each repo
		}
		fmt.Println("getting release info for repo", repoScout.RepoName)
		// Fetch release info for the repo
		releaseInfo, err := dao.ServiceRepo.BuildService.GetReleaseInfo([]string{repoScout.RepoName}, 1)
		fmt.Println("release info is ", releaseInfo, err)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Error fetching release info",
				"error":   err.Error(),
			})
			return
		}

		if len(releaseInfo) > 0 {
			repoResponse["release_info"] = releaseInfo[0]
			repoResponse["latest_image_url"] = repoScout.DockerBaseURL + ":" + releaseInfo[0].Releases.TagName
		}
		// Fetch deployments for the repo
		for _, deployment := range repoScout.Deployments {
			deploymentInfo, err := dao.ServiceRepo.DeploymentService.GetDeploymentByName(request.Namespace, deployment)
			if err != nil {
				// Log error and continue to next deployment
				fmt.Println("Error fetching deployment info:", err)
				logger.Logger.Error("Error GetDeploymentByName", zap.Any("err:", err.Error()))
				continue
			}
			if deploymentInfo != nil && utils.GetDockertagFromURL(deploymentInfo.Image) != releaseInfo[0].Releases.TagName {
				deploymentInfo.OutOfSync = true
			}
			// Append deployment data to the repo response
			repoResponse["deployments"] = append(repoResponse["deployments"].([]map[string]interface{}), map[string]interface{}{
				"deployment_info": deploymentInfo,
			})
		}

		// Append the repo data to the response list
		response = append(response, repoResponse)
	}

	// Return the final response
	ctx.JSON(http.StatusOK, response)
}
