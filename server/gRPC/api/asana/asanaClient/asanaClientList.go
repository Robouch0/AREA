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
				Name:    "Create a Board on Asana",
				RefName: "createBoard",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"boardName": {
						Value:       "",
						Type:        "string",
						Description: "The name of the board you want to create",
						Required:    true,
					},
					"boardDescription": {
						Value:       "",
						Type:        "string",
						Description: "A small description of the board",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}
