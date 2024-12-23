package svc

import (
	adapter "deployment-service/apps/repository/adapter"
	"deployment-service/constants"
	"deployment-service/logger"
	model_event_logger "deployment-service/models/model.eventlogger"
	"time"

	"go.uber.org/zap"
)

type EventLoggerService struct {
	repository *adapter.Repository
}

func (svc EventLoggerService) LogEvent(payload model_event_logger.Event) ([]map[string]interface{}, error) {
	logger.EventLogger.Info(constants.SERVICE_NAME, zap.Any("payload", payload))
	deployments, err := svc.repository.Kubernetes.ListDeployments("default")
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
