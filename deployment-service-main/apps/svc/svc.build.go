package svc

import (
	"context"
	adapter "deployment-service/apps/repository/adapter"
	"deployment-service/logger"
	model_build "deployment-service/models/model.build"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"go.uber.org/zap"
)

type BuildService struct {
	repository *adapter.Repository
}

func (svc BuildService) CreateNewRepoScout(payload model_build.RepoScout) (map[string]interface{}, error) {
	fmt.Println("payload is ", payload)
	result, err := svc.repository.MongoDB.InsertOne("REPO_SCOUTS", payload)
	if err != nil {
		logger.Logger.Error("Error while inserting new repo scout", zap.Any(logger.KEY_ERROR, err.Error()))
	}
	return map[string]interface{}{
		"result": payload,
		"data":   result,
	}, nil
}

func (svc BuildService) GetAllRepoScouts(namespace string) ([]model_build.RepoScout, error) {
	// Prepare the query filter to match the specified namespace
	filter := bson.D{
		{"namespace", namespace},
	}

	// Fetch all documents that match the filter from the REPO_SCOUTS collection
	cursor, err := svc.repository.MongoDB.GetAll("REPO_SCOUTS", filter)
	defer cursor.Close(context.TODO()) // Close the cursor after we're done

	var result = []model_build.RepoScout{}
	for cursor.Next(context.TODO()) {
		var repoScout model_build.RepoScout
		if err := cursor.Decode(&repoScout); err != nil {
			fmt.Println("Error decoding document:", err)
			return nil, err
		}
		result = append(result, repoScout)
	}

	// Check for cursor iteration errors
	if err := cursor.Err(); err != nil {
		fmt.Println("Error iterating over cursor:", err)
		return nil, err
	}

	if err != nil {
		// Log the error if fetching fails
		logger.Logger.Error("Error while fetching all repo scouts", zap.Any(logger.KEY_ERROR, err.Error()))
		return nil, err
	}

	// Return the fetched result
	return result, nil
}

func (svc BuildService) GetReleaseInfo(repoURLs []string, topK int) ([]model_build.RepoReleases, error) {
	var result []model_build.RepoReleases

	// Iterate over each repository URL
	for _, repoURL := range repoURLs {
		// Build the GitHub API URL for releases
		apiURL := fmt.Sprintf("https://api.github.com/repos/%s/releases/latest", repoURL)
		// Send HTTP GET request to GitHub API
		resp, err := http.Get(apiURL)
		if err != nil {
			fmt.Println("error is 80", err)
			continue
			// return nil, fmt.Errorf("failed to fetch releases for %s: %w", repoURL, err)
		}
		defer resp.Body.Close()

		// Check the HTTP status code
		if resp.StatusCode != 200 {
			fmt.Println("the error is ", resp)
			fmt.Println("error message is ", resp.Body)
			continue
			// return nil, fmt.Errorf("unexpected status code %d for %s and %s", resp.StatusCode, repoURL, resp)
		}

		// Log the body for debugging
		body1, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("error is 97", err)
			continue

			// return nil, fmt.Errorf("failed to read response body for %s: %w", repoURL, err)
		}

		// Decode the response
		var releases model_build.ReleaseInfo
		if err := json.Unmarshal(body1, &releases); err != nil {
			fmt.Println("error is 105", err)
			continue
			// return nil, fmt.Errorf("failed to decode response for %s: %w", repoURL, err)
		}

		// // Select the top K releases
		// var topReleases []model_build.ReleaseInfo
		// for i, release := range releases {
		// 	if i >= topK {
		// 		break
		// 	}

		// 	// Add the release to the topReleases slice
		// 	topReleases = append(topReleases, release)
		// }

		// Append to the result with the formatted repo info
		result = append(result, model_build.RepoReleases{
			RepoURL:  "https://github.com/" + repoURL, // Keep the repository URL for reference
			Releases: releases,                        // Append the top releases for the repo
		})
	}

	return result, nil
}
