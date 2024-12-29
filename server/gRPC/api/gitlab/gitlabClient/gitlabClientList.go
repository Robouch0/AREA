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

		Microservices: []IServ.MicroserviceDescriptor{},
	}
	return status, nil
}
