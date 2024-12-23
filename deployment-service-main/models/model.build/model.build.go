package model_build

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepoScout struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	RepoURL       string             `bson:"github_url" json:"github_url"`
	RepoName      string             `bson:"repo_name" json:"repo_name"`
	DockerBaseURL string             `bson:"docker_base_url" json:"docker_base_url"`
	Namespace     string             `bson:"namespace" json:"namespace"`
	Deployments   []string           `bson:"deployments" json:"deployments"`
	CreatedAt     time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt     time.Time          `bson:"updatedAt,omitempty" json:"updatedAt"`
}

type ReleaseInfo struct {
	HtmlURL     string    `json:"html_url"`
	TagName     string    `json:"tag_name"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
}

type GetRepoScoutResp struct {
	RepoName            string `json:"repo_name"`
	RepoUrl             string `json:"repo_url"`
	LatestReleaseTag    string `json:"latest_release_tag"`
	LatestReleaseTagUrl string `json:"latest_release_tag_url"`
	DeploymentStatus    string `json:"deployment_status"`
	DeployedVersion     string `json:"deployed_version"`
	DeploymentInfoUrl   string `json:"deployment_info_url"`
	ActionMessage       string `json:"action_message"`
}

type TenantResourceResp struct {
	KubernetesVersion string `json:"kubernetes_version"`
	NoOfPods          int64  `json:"no_of_pods"`
	NoOfDeployments   int64  `json:"no_of_deployments"`
	NoOfServices      int64  `json:"no_of_services"`
	TenantUsername    string `json:"tenant_username"`
}

// RepoReleases holds information about a repository and its top releases
type RepoReleases struct {
	RepoURL  string      `json:"repo_url"`
	Releases ReleaseInfo `json:"releases"`
}
