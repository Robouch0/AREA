//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFaceClientList
//

package huggingFace

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
		},
	}
	return status, nil
}
