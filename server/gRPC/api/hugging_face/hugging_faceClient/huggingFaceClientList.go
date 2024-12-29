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

		Microservices: []IServ.MicroserviceStatus{
			{
				Name:    "Text Generation",
				RefName: "textGen",
				Type:    "reaction",

				Ingredients: map[string]string{
					"model":  "string",
					"inputs": "string",
				},
			},
			{
				Name:    "Create repository",
				RefName: "createRepo",
				Type:    "reaction",

				Ingredients: map[string]string{
					"type":         "string", // Default is model
					"name":         "string",
					"organisation": "string", // optionnal
					"private":      "string", // optionnal
				},
			},
			// {
			// 	Name:    "Check if someone push in the repository",
			// 	RefName: "push",
			// 	Type:    "action",

			// 	Ingredients: map[string]string{
			// 		"type": "string",
			// 		"name": "string",
			// 	},
			// },
			// {
			// 	Name:    "Check if someone create a pull request",
			// 	RefName: "pr",
			// 	Type:    "action",

			// 	Ingredients: map[string]string{
			// 		"type": "string",
			// 		"name": "string",
			// 	},
			// },
			// {
			// 	Name:    "Check if someone create a discussion",
			// 	RefName: "discussion",
			// 	Type:    "action",

			// 	Ingredients: map[string]string{
			// 		"type": "string",
			// 		"name": "string",
			// 	},
			// },
		},
	}
	return status, nil
}
