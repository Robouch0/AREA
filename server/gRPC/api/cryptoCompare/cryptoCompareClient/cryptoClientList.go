//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// cryptoClientList
//

package cryptocompareclient

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (crypto *CryptoClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Crypto",
		RefName: "crypto",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Each hour, when crypto currency value exceed a certain amount trigger happens",
				RefName: "cryptoExceed",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{
					"crypto_currency": {
						Value:       "",
						Type:        "string",
						Description: "Crypto currency that will be checked",
						Required:    true,
					},
					"currency": {
						Value:       "",
						Type:        "string",
						Description: "currency value that will be tracked",
						Required:    true,
					},
					"threshold": {
						Value:       "",
						Type:        "int",
						Description: "Value of the currency that will trigger the action",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO: Matthieu
			},
			{
				Name:    "Each hour, when crypto currency value is lower a certain amount trigger happens",
				RefName: "cryptoLower",
				Type:    "action",
				Ingredients: map[string]IServ.IngredientDescriptor{
					"crypto_currency": {
						Value:       "",
						Type:        "string",
						Description: "Crypto currency that will be checked",
						Required:    true,
					},
					"currency": {
						Value:       "",
						Type:        "string",
						Description: "currency value that will be tracked",
						Required:    true,
					},
					"threshold": {
						Value:       "",
						Type:        "int",
						Description: "Value of the currency that will trigger the action",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO: Matthieu
			},
		},
	}
	return status, nil
}
