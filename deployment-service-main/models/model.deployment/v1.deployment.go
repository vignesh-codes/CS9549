package model_deployment

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateDeploymentRequest struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Namespace     string             `bson:"namespace" json:"namespace"`
	ContainerPort int32              `bson:"containerPort" json:"container_port"`
	Image         string             `bson:"image" json:"image"`
	Replicas      int32              `bson:"replicas" json:"replicas"`
	RepoScoutId   string             `bson:"repo_scout_id" json:"repo_scout_id"`
	Status        string             `bson:"status" json:"status"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type UpdateDeploymentReq struct {
	Name     string `json:"name"`
	Replicas int32  `json:"replicas"`
	Image    string `json:"image"`
}

type DeploymentInfo struct {
	DeploymentName    string                 `json:"deployment_name"`
	Age               string                 `json:"age"`
	Status            string                 `json:"status"`
	DesiredReplicas   int                    `json:"desired_replicas"`
	CurrentReplicas   int                    `json:"current_replicas"`
	Image             string                 `json:"image"`
	AvailableReplicas int                    `json:"available_replicas"`
	OtherInfo         map[string]interface{} `json:"other_info"`
	OutOfSync         bool                   `json:"out_of_sync"`
}
