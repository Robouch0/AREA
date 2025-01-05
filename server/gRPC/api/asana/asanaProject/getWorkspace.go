//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana utils project
//

package asana_project

import (
	asana_generics "area/gRPC/api/asana/asanaGenerics"
	"area/utils"
	http_utils "area/utils/httpUtils"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type WorkspaceInfo struct {
	Name          string `json:"name,omitempty"`
	Gid           string `json:"gid,omitempty"`
	RessourceType string `json:"ressource_type,omitempty"`
}

func GetGidByWorkspace(workSpaceName string, list *asana_generics.AsanaBaseBody[[]WorkspaceInfo]) (string, error) {
	for _, workspaceInfo := range list.Data {
		if workspaceInfo.Name == workSpaceName {
			return workspaceInfo.Gid, nil
		}
	}
	return "", status.Errorf(codes.NotFound, "Workspace name not found in the list")
}

func ListWorkspace(accesToken string) (*asana_generics.AsanaBaseBody[[]WorkspaceInfo], error) {
	url := "https://app.asana.com/api/1.0/workspaces"

	postRequest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating API request: %v", err)
	}

	postRequest.Header.Set("Accept", "application/json")
	postRequest.Header.Set("Authorization", "Bearer "+accesToken)

	res, err := http_utils.SendHttpRequest(postRequest, 200)
	if err != nil {
		return nil, err
	}

	list, err := utils.IoReaderToStruct[asana_generics.AsanaBaseBody[[]WorkspaceInfo]](&res.Body)
	if err != nil {
		return nil, err
	}
	return list, nil
}
