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
					"private":      "bool",   // optionnal
				},
			},
			{
				Name:    "Move repository",
				RefName: "moveRepo",
				Type:    "reaction",

				Ingredients: map[string]string{
					"fromRepo": "string",
					"toRepo":   "string",
					"type":     "string",
				},
			},
		},
	}
	return status, nil
}
