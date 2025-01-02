//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceClientList
//

package huggingFace_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (hfCli *HuggingFaceServiceClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Hugging Face",
		RefName: "hf",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Text Generation",
				RefName: "textGen",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"model": {
						Value:       "",
						Type:        "string",
						Description: "Name of the model",
						Required:    true,
					},
					"inputs": {
						Value:       "",
						Type:        "string",
						Description: "Input phrase of the model",
						Required:    true,
					},
				},
			},
			{
				Name:    "Create repository",
				RefName: "createRepo",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"type": {
						Value:       "model",
						Type:        "string",
						Description: "Type of repository",
						Required:    true,
					},
					"name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
					"organisation": {
						Value:       "",
						Type:        "string",
						Description: "Name of the origanisation",
						Required:    false,
					},
					"private": {
						Value:       false,
						Type:        "bool",
						Description: "Whether or not the repository is private",
						Required:    false,
					},
				},
			},
			{
				Name:    "Move repository",
				RefName: "moveRepo",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"fromRepo": {
						Value:       "",
						Type:        "string",
						Description: "Repository name that will be moved",
						Required:    true,
					},
					"toRepo": {
						Value:       "",
						Type:        "string",
						Description: "Destination of the repository",
						Required:    true,
					},
					"type": {
						Value:       "model",
						Type:        "string",
						Description: "Type of repository",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}
