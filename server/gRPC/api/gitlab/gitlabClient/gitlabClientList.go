//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gitlabClientList
//

package gitlab_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (git *GitlabClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Gitlab",
		RefName: "gitlab",

		Microservices: []IServ.MicroserviceStatus{
			IServ.MicroserviceStatus{
				Name:    "Create a file on a repository",
				RefName: "createFile",
				Type:    "reaction",

				Ingredients: map[string]string{
					"file_path" : "string",
					"id" : "string",
					"branch": "string",
					"commit_message":  "string",
					"content" : "string",
				},
			},
		},
	}
	return status, nil
}
