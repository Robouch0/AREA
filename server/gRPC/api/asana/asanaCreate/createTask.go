//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana utils project
//

package asana_create

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

type CreateTaskData struct {
	Name      string   `json:"name,omitempty"`
	Notes     string   `json:"notes,omitempty"`
	Projects  []string `json:"projects,omitempty"`
	Completed bool     `json:"completed,omitempty"`
	DueOn     string   `json:"due_on,omitempty"`
}

func CreateTask(accessToken string, body *asana_generics.AsanaBaseBody[CreateTaskData]) (*asana_generics.AsanaBaseBody[CreateTaskData], error) {
	url := "https://app.asana.com/api/1.0/tasks"

	b, err := json.Marshal(body)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error when Marshalizing the request body: %v", err)
	}

	postRequest, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("error creating API request: %v", err)
	}

	postRequest.Header.Set("Accept", "application/json")
	postRequest.Header.Set("Content-Type", "application/json")
	postRequest.Header.Set("Authorization", "Bearer "+accessToken)

	_, err = http_utils.SendHttpRequest(postRequest, 201)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "error when sending request: %v", err)
	}

	return body, nil
}
