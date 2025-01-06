//
// EPITECH PROJECT, 2024
// AREA
// File description:
// asana Board service
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

type CreateSectionData struct {
	Name       string `json:"name,omitempty"`
	ProjectGid string `json:"-"`
}

func CreateSection(accessToken string, body *asana_generics.AsanaBaseBody[CreateSectionData]) (*asana_generics.AsanaBaseBody[CreateSectionData], error) {
	url := fmt.Sprintf("https://app.asana.com/api/1.0/projects/%s/sections", body.Data.ProjectGid)

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
