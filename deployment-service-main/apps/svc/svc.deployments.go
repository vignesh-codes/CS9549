package svc

import (
	"context"
	adapter "deployment-service/apps/repository/adapter"
	"deployment-service/logger"
	model_build "deployment-service/models/model.build"
	model_deployment "deployment-service/models/model.deployment"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.uber.org/zap"
)

type DeploymentService struct {
	repository *adapter.Repository
}

func (svc DeploymentService) GetDeploymentsByNamespace(namespace string) ([]map[string]interface{}, error) {
	deployments, err := svc.repository.Kubernetes.ListDeployments(namespace)
	var deploymentInfo []map[string]interface{}
	if err == nil {

		for _, d := range deployments {
			replicasNeeded := int32(0)
			replicasAvailable := int32(0)
			if d.Spec.Replicas != nil {
				replicasNeeded = *d.Spec.Replicas
			}
			if d.Status.AvailableReplicas > 0 {
				replicasAvailable = d.Status.AvailableReplicas
			}

			// Calculate the age of the deployment
			age := time.Since(d.CreationTimestamp.Time)

			// Fetch resource utilization if you have metrics access (optional, requires Metrics API)
			// For example, get memory and CPU usage per pod (implement separately)
			// resourceUtilization := getDeploymentResourceUtilization(d.Name, namespace, k)

			// Create YAML manifest string
			// deploymentYAML, err := yaml.Marshal(d)
			// if err != nil {
			// 	fmt.Printf("failed to marshal deployment %s to YAML: %v\n", d.Name, err)
			// 	deploymentYAML = []byte("Error generating YAML")
			// }

			deploymentDetails := map[string]interface{}{
				"Deployment Name":    d.Name,
				"Namespace":          d.Namespace,
				"Replicas Needed":    replicasNeeded,
				"Replicas Available": replicasAvailable,
				"Status":             d.Status.Conditions,
				"Age":                age.String(),
				"Creation Time":      d.CreationTimestamp.Time,
				// "Resource Utilization": resourceUtilization, // CPU/Memory data as needed
				// "Deployment Manifest": string(deploymentYAML),
			}

			deploymentInfo = append(deploymentInfo, deploymentDetails)
		}
	}
	// fmt.Println(deploymentInfo)
	// svc.repository.Kubernetes.ListServices("default")
	return deploymentInfo, nil
}

func (svc DeploymentService) GetTenantKubernetesInfo(namespace string) (model_build.TenantResourceResp, error) {
	var resp = model_build.TenantResourceResp{}
	// Create Namespace if not exists
	nserr := svc.repository.Kubernetes.CreateNamespaceIfNotExists(namespace)
	if nserr != nil {
		fmt.Println("Error creating namespace %s: %v", namespace, nserr)
		return resp, nserr
	}
	// Fetch Pods
	pods, err := svc.repository.Kubernetes.ListPods(namespace)
	if err != nil {
		return resp, err
	}

	// Fetch Kubernetes Version
	k8sVersion, err := svc.repository.Kubernetes.GetKubernetesVersion()
	if err != nil {
		return resp, err
	}

	// Fetch Deployments
	deployments, err := svc.repository.Kubernetes.ListDeployments(namespace)
	if err != nil {
		return resp, err
	}

	// Fetch Services
	services, err := svc.repository.Kubernetes.ListServices(namespace)
	if err != nil {
		return resp, err
	}

	resp.KubernetesVersion = k8sVersion
	resp.NoOfDeployments = int64(len(deployments))
	resp.NoOfPods = int64(len(pods))
	resp.NoOfServices = int64(len(services))
	resp.TenantUsername = namespace
	return resp, nil
}

// DeploymentService: Updated method to map available replicas and status correctly
func (svc DeploymentService) GetDeploymentByName(namespace, deploymentName string) (*model_deployment.DeploymentInfo, error) {
	// Get structured deployment data
	kubernetesManifest, err := svc.repository.Kubernetes.GetDeploymentByName(namespace, deploymentName)
	if err != nil {
		return nil, err
	}

	// Retrieve replicas info
	desiredReplicas := int(kubernetesManifest.DesiredReplicas)
	currentReplicas := int(kubernetesManifest.CurrentReplicas)
	availableReplicas := int(kubernetesManifest.AvailableReplicas) // Add available replicas
	svcInfo, err := svc.repository.Kubernetes.GetServiceInfo(namespace, deploymentName+"-service")
	if err != nil {
		svcInfo = "UNDEFINED"
	}
	// Populate DeploymentInfo struct
	deploymentInfo := &model_deployment.DeploymentInfo{
		DeploymentName:    deploymentName,
		Age:               kubernetesManifest.Age,
		Status:            kubernetesManifest.Status, // Adjust to include available status like "Available", "Progressing", etc.
		DesiredReplicas:   desiredReplicas,
		CurrentReplicas:   currentReplicas,
		Image:             kubernetesManifest.Image,
		AvailableReplicas: availableReplicas, // Add available replicas field
		OtherInfo: map[string]interface{}{
			"kuberenetes_spec": kubernetesManifest.Spec,
			"endpoint":         svcInfo,
		},
	}

	return deploymentInfo, nil
}

func (svc DeploymentService) GetLatestEvents(namespace string, topK int) ([]string, error) {
	return svc.repository.Kubernetes.GetLatestEvents(namespace, topK)
}

// UpdateDeploymentReplicas updates the number of replicas for a given deployment in Kubernetes
// and updates the corresponding MongoDB document.
func (svc DeploymentService) UpdateDeploymentByName(namespace, deploymentName string, image string, replicas int32) (map[string]interface{}, error) {
	// Retrieve the current deployment object
	fmt.Println("143 ---- ", deploymentName, replicas, image)

	deployment, err := svc.GetDeploymentFromDBByName(namespace, deploymentName)
	if err != nil {
		return nil, fmt.Errorf("failed to get deployment %s in namespace %s: %w", deploymentName, namespace, err)
	}
	fmt.Println("148 ---- ", deployment.Replicas, deployment.Image)
	needToUpdateDb := false
	if replicas == -1 {
		replicas = deployment.Replicas

	}
	if image == "" {
		image = deployment.Image

	}

	// Update the replicas in Kubernetes deployment
	if replicas != deployment.Replicas || image != deployment.Image {
		needToUpdateDb = true
		err := svc.repository.Kubernetes.UpdateDeploymentReplicasAndImage(namespace, deploymentName, replicas, image)
		if err != nil {
			return nil, fmt.Errorf("failed to update replicas in Kubernetes: %w", err)
		}
	}
	if !needToUpdateDb {
		return map[string]interface{}{
			"message": fmt.Sprintf("Successfully updated replicas to %d and image to %s for deployment %s in Kubernetes", replicas, deploymentName, image),
		}, nil
	}
	// Update the corresponding MongoDB document
	resp, err := svc.updateDeploymentInMongoDB(deploymentName, image, replicas)
	if err != nil {
		return nil, fmt.Errorf("failed to update MongoDB for deployment %s: %w", deploymentName, err)
	}

	fmt.Printf("Successfully updated replicas to %d and image to %s for deployment %s in Kubernetes", replicas, deploymentName, image)
	return resp, nil
}

// updateDeploymentInMongoDB updates the replica count for the deployment in MongoDB's DEPLOYMENTS collection.
func (svc DeploymentService) updateDeploymentInMongoDB(deploymentName string, image string, replicas int32) (map[string]interface{}, error) {
	// Construct the filter and update for MongoDB
	fmt.Println("updating this item ", deploymentName, image, replicas)
	filter := bson.M{"name": deploymentName}
	update := bson.M{
		"$set": bson.M{
			"image":    image,
			"replicas": replicas,
		},
	}

	// Update the MongoDB document
	res, err := svc.repository.MongoDB.UpdateOne("DEPLOYMENTS", filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to update MongoDB document: %w", err)
	}

	return map[string]interface{}{
		"result": res,
	}, nil
}

func (svc DeploymentService) CreateNamespaceIfNotExists(namespace string) error {
	return svc.repository.Kubernetes.CreateNamespaceIfNotExists(namespace)
}

func (svc DeploymentService) CreateDeployment(payload *model_deployment.CreateDeploymentRequest) (interface{}, error) {
	// check if build exists
	var result bson.M
	objectId, err := primitive.ObjectIDFromHex(payload.RepoScoutId)
	if err != nil {
		return nil, errors.New("Invalid RepoScoutId format")
	}
	err1 := svc.repository.MongoDB.FindOne("REPO_SCOUTS", bson.M{"_id": objectId}).Decode(&result)
	if err1 != nil {
		fmt.Printf("FindOne error: %v\n", err1)
		return nil, errors.New("Repo scout id not found")
	}
	// Create the Deployment
	err = svc.repository.Kubernetes.CreateDeployment(payload.Namespace, payload.Name,
		payload.Image, payload.Replicas, payload.ContainerPort, "50m", "0.2Gi")
	if err != nil {
		return nil, fmt.Errorf("failed to create deployment: %w", err)
	}

	err = svc.repository.Kubernetes.CreateService(payload.Namespace, payload.Name+"-service",
		payload.Name, 80, payload.ContainerPort)
	if err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}
	payload.Status = "ACTIVE"

	// Insert the deployment data into MongoDB
	_, err = svc.repository.MongoDB.InsertOne("DEPLOYMENTS", payload)
	if err != nil {
		logger.Logger.Error("Error while inserting new deployment", zap.Any(logger.KEY_ERROR, err.Error()))
		return nil, err
	}
	fmt.Println("repo scout id is ", payload.RepoScoutId)
	// update repo scout based on RepoScoutId from payload
	repo_scout_id, err := primitive.ObjectIDFromHex(payload.RepoScoutId)
	fmt.Println("repo scout is is , ", repo_scout_id, err)
	filter := bson.M{"_id": repo_scout_id}

	// Define the update operation to push the new DeploymentID to the Deployments array
	update := bson.M{
		"$push": bson.M{"deployments": payload.Name},
		"$set":  bson.M{"updatedAt": time.Now()},
	}

	// Perform the update operation
	scout_repo_result, err := svc.repository.MongoDB.UpdateOne("REPO_SCOUTS", filter, update)
	if err != nil {
		logger.Logger.Error("Error while updating repo scouts", zap.Any(logger.KEY_ERROR, err.Error()))
		return nil, err
	}

	// Log success if document was updated
	if scout_repo_result.ModifiedCount > 0 {
		logger.Logger.Info("RepoScout updated successfully", zap.Any("RepoScoutId", payload.ID))
	} else {
		logger.Logger.Warn("No RepoScout document found with specified RepoScoutId", zap.Any("RepoScoutId", payload.ID))
	}
	return payload, nil
}

func (svc DeploymentService) GetAllDeploymentsFromDBByNamespace(namespace string) ([]model_deployment.CreateDeploymentRequest, error) {

	filter := bson.M{"namespace": namespace}
	var results = []model_deployment.CreateDeploymentRequest{}
	cursor, err := svc.repository.MongoDB.FindMany("DEPLOYMENTS", filter)
	if err != nil {
		fmt.Println("error in fetchning many deployments", err)
		return nil, err
	}
	// Iterate over the cursor to decode each document into the slice
	for cursor.Next(context.TODO()) {
		var deployment model_deployment.CreateDeploymentRequest
		if err := cursor.Decode(&deployment); err != nil {
			return nil, fmt.Errorf("error decoding document: %w", err)
		}
		results = append(results, deployment)
	}

	// Check if there were any errors during iteration
	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over cursor: %w", err)
	}
	fmt.Println("res is ", results)
	return results, nil
}

func (svc DeploymentService) GetDeploymentFromDBByName(namespace, deplyomentName string) (*model_deployment.CreateDeploymentRequest, error) {
	filter := bson.M{"namespace": namespace, "name": deplyomentName}
	var result = model_deployment.CreateDeploymentRequest{}
	res := svc.repository.MongoDB.FindOne("DEPLOYMENTS", filter)
	if err := res.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (svc DeploymentService) DeleteDeployment(namespace, deploymentName, RepoScoutId string) (map[string]interface{}, error) {
	// Delete the Deployment from Kubernetes
	err := svc.repository.Kubernetes.DeleteDeployment(namespace, deploymentName)
	if err != nil {
		// return nil, fmt.Errorf("failed to delete deployment: %w", err)
	}

	// Delete the associated Service from Kubernetes
	err = svc.repository.Kubernetes.DeleteService(namespace, deploymentName+"-service")
	if err != nil {
		return nil, fmt.Errorf("failed to delete service: %w", err)
	}

	// Remove the deployment record from MongoDB
	filter := bson.M{"namespace": namespace, "name": deploymentName}
	deploymentDeleteResult, err := svc.repository.MongoDB.DeleteOne("DEPLOYMENTS", filter)
	if err != nil {
		logger.Logger.Error("Error while deleting deployment from MongoDB", zap.Any(logger.KEY_ERROR, err.Error()))
		return nil, err
	}

	// Update the RepoScout document to remove the deployment reference
	repo_scout_id, err := primitive.ObjectIDFromHex(RepoScoutId)
	fmt.Println("repo scout is is , ", repo_scout_id, err)
	updateFilter := bson.M{"_id": repo_scout_id}
	update := bson.M{
		"$pull": bson.M{"deployments": deploymentName},
		"$set":  bson.M{"updatedAt": time.Now()},
	}

	repoScoutUpdateResult, err := svc.repository.MongoDB.UpdateOne("REPO_SCOUTS", updateFilter, update)
	if err != nil {
		logger.Logger.Error("Error while updating RepoScout in MongoDB", zap.Any(logger.KEY_ERROR, err.Error()))
		return nil, err
	}

	// Log success if the RepoScout document was updated
	if repoScoutUpdateResult.ModifiedCount > 0 {
		logger.Logger.Info("RepoScout updated successfully", zap.String("RepoScoutId", RepoScoutId))
	} else {
		logger.Logger.Warn("No RepoScout document found with specified RepoScoutId", zap.String("RepoScoutId", RepoScoutId))
	}

	// Return details of the delete operation
	return map[string]interface{}{
		"deployment_delete_result": deploymentDeleteResult,
		"repo_scout_update_result": repoScoutUpdateResult,
	}, nil
}
