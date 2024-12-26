//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleList
//

package google_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (google *GoogleClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Google",
		RefName: "google",

		Microservices: []IServ.MicroserviceStatus{
			IServ.MicroserviceStatus{
				Name:    "Send an email to a specific user",
				RefName: "gmail/sendEmailMe",
				Type:    "reaction",

				Ingredients: map[string]string{
					"to":           "string",
					"subject":      "string",
					"body_message": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Delete an email of a specific user",
				RefName: "gmail/deleteEmailMe",
				Type:    "reaction",

				Ingredients: map[string]string{
					"subject": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Move an email to trash",
				RefName: "gmail/moveToTrash",
				Type:    "reaction",

				Ingredients: map[string]string{
					"subject": "string",
				},
			},
			IServ.MicroserviceStatus{
				Name:    "Move an email from trash",
				RefName: "gmail/moveFromTrash",
				Type:    "reaction",

				Ingredients: map[string]string{
					"subject": "string",
				},
			},
		},
	}
	return status, nil
}
