//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana utils project
//

package asana_project

import (
	asana_generics "area/gRPC/api/asana/asanaGenerics"
	http_utils "area/utils/httpUtils"
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type CreateProjectData struct {
	Name         string `json:"name,omitempty"`
	Color        string `json:"color,omitempty"`
	DefaultView  string `json:"default_view,omitempty"`
	WorkspaceGid string `json:"workspace,omitempty"`
}

func CreateProject(accesToken string, body *asana_generics.AsanaBaseBody[CreateProjectData]) (*asana_generics.AsanaBaseBody[CreateProjectData], error) {
	url := "https://app.asana.com/api/1.0/projects"

	b, err := json.Marshal(body)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error when Marshalizing the reques body: %v", err)
	}
	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("error creating API request: %v", err)
	}

	postRequest.Header.Set("Accept", "application/json")
	postRequest.Header.Set("Content-Type", "application/json")
	postRequest.Header.Set("Authorization", "Bearer "+accesToken)

	_, err = http_utils.SendHttpRequest(postRequest, 201)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error when sending request : %v", err)
	}
	return body, nil
}
