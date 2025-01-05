//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trelloClientList
//

package asana_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (asana *AsanaClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Asana",
		RefName: "asana",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Create a Project on Asana",
				RefName: "createProject",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"projectName": {
						Value:       "",
						Type:        "string",
						Description: "The name of the project you want to create",
						Required:    true,
					},
					"color": {
						Value:       "",
						Type:        "string",
						Description: "The background color theme of your project",
						Required:    true,
					},
					"defaultView": {
						Value:       "",
						Type:        "string",
						Description: "The default view apply to your project (ex: board)",
						Required:    true,
					},
					"workspaceName": {
						Value:       "",
						Type:        "string",
						Description: "The targeted workspace name",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}
