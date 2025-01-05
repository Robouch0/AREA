//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miroClientList
//

package miro_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (miro *MiroClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Miro",
		RefName: "miro",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Each time a new sticker is added a trigger is activated",
				RefName: "watchItemCreated",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"boardId": {
						Value:       "",
						Type:        "string",
						Description: "Identifier of the board",
						Required:    true,
					},
					"status": {
						Value:       "",
						Type:        "bool",
						Description: "Wheter or not the webhook is activated",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}
