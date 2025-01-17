//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana utils project
//

package asana_get

import (
	asana_generics "area/gRPC/api/asana/asanaGenerics"
	conv_utils "area/utils/convUtils"
	http_utils "area/utils/httpUtils"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProjectInfo struct {
	Name          string `json:"name,omitempty"`
	Gid           string `json:"gid,omitempty"`
	RessourceType string `json:"resource_type,omitempty"`
}

func GetGidByProject(projectName string, list *asana_generics.AsanaBaseBody[[]ProjectInfo]) (string, error) {
	for _, projectInfo := range list.Data {
		if projectInfo.Name == projectName {
			return projectInfo.Gid, nil
		}
	}
	return "", status.Errorf(codes.NotFound, "Project name not found in the list")
}

// ListAllProjects maybe we should later filter with a workspace name using another endpoint
func ListAllProjects(accessToken string) (*asana_generics.AsanaBaseBody[[]ProjectInfo], error) {
	url := "https://app.asana.com/api/1.0/projects"

	postRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating API request: %v", err)
	}

	postRequest.Header.Set("Accept", "application/json")
	postRequest.Header.Set("Authorization", "Bearer "+accessToken)

	res, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}

	list, err := conv_utils.IoReaderToStruct[asana_generics.AsanaBaseBody[[]ProjectInfo]](&res.Body)
	if err != nil {
		return nil, err
	}
	return list, nil
}
