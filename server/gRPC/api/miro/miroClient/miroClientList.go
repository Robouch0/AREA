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
				Name:    "Create a sticky note in miro board",
				RefName: "createStickyNotes",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"boardId": {
						Value:       "",
						Type:        "string",
						Description: "ID of the board where the reaction happens",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "Content of the sticky note",
						Required:    true,
					},
					"shape": {
						Value:       "",
						Type:        "string",
						Description: "Shape of the sticky note rectangle or square",
						Required:    true,
					},
				},
			},
			{
				Name:    "Create a text note in miro board",
				RefName: "createTextNotes",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"boardId": {
						Value:       "",
						Type:        "string",
						Description: "ID of the board where the reaction happens",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "Content of the text note",
						Required:    true,
					},
				},
			},
			{
				Name:    "Create a card note in miro board",
				RefName: "createCardNotes",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"boardId": {
						Value:       "",
						Type:        "string",
						Description: "ID of the board where the reaction happens",
						Required:    true,
					},
					"description": {
						Value:       "",
						Type:        "string",
						Description: "A short text description to add context about the card",
						Required:    true,
					},
					"title": {
						Value:       "",
						Type:        "string",
						Description: "A short text header for the card",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}
