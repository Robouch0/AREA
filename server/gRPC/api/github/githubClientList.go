//
// EPITECH PROJECT, 2024
// AREA [WSL: Ubuntu]
// File description:
// githubClientList
//

package github

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (git *GithubClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Github",
		RefName: "github",

		Microservices: []IServ.MicroserviceStatus{
			IServ.MicroserviceStatus{
				Name:    "Trigger every new push of a repository",
				RefName: "triggerPush",
				Type:    "action",

				Ingredients: map[string]string{
					"owner": "string",
					"repo":  "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Update Repository Informations",
				RefName: "updateRepo",
				Type:    "reaction",

				Ingredients: map[string]string{
					"owner":       "string",
					"repo":        "string",
					"name":        "string",
					"description": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Update a file in a repository",
				RefName: "updateFile",
				Type:    "reaction",

				Ingredients: map[string]string{
					"owner":   "string",
					"repo":    "string",
					"path":    "string",
					"message": "string",
					"content": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Delete Repository File",
				RefName: "deleteFile",
				Type:    "reaction",

				Ingredients: map[string]string{
					"owner":   "string",
					"repo":    "string",
					"path":    "string",
					"message": "string",
				},
			},
		},
	}
	return status, nil
}