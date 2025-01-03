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

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Send an email to a specific user",
				RefName: "gmail/sendEmailMe",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"to": {
						Value:       "",
						Type:        "string",
						Description: "Destination e-mail",
						Required:    true,
					},
					"subject": {
						Value:       "",
						Type:        "string",
						Description: "Subject of the email",
						Required:    true,
					},
					"body_message": {
						Value:       "",
						Type:        "string",
						Description: "Message Inside the email",
						Required:    true,
					},
				},
			},
			{
				Name:    "Delete an email of a specific user",
				RefName: "gmail/deleteEmailMe",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"subject": {
						Value:       "",
						Type:        "string",
						Description: "Subject of the email",
						Required:    true,
					},
				},
			},
			{
				Name:    "Move an email to trash",
				RefName: "gmail/moveToTrash",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"subject": {
						Value:       "",
						Type:        "string",
						Description: "Subject of the email",
						Required:    true,
					},
				},
			},
			{
				Name:    "Move an email from trash",
				RefName: "gmail/moveFromTrash",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"subject": {
						Value:       "",
						Type:        "string",
						Description: "Subject of the email",
						Required:    true,
					},
				},
			},
			{
				Name:    "Create a new label",
				RefName: "gmail/createLabel",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the label",
						Required:    true,
					},
					"messageListVisibility": {
						Value:       "",
						Type:        "string",
						Description: "Visibility of the message with this label",
						Required:    true,
					},
					"LabelListVisibility": {
						Value:       "",
						Type:        "string",
						Description: "Visibility of the label",
						Required:    true,
					},
					"type": {
						Value:       "",
						Type:        "string",
						Description: "Type of the label (User/System)",
						Required:    true,
					},
				},
			},
			{
				Name:    "Update a label",
				RefName: "gmail/updateLabel",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"old_name": {
						Value:       "",
						Type:        "string",
						Description: "Old name of the label",
						Required:    true,
					},
					"new_name": {
						Value:       "",
						Type:        "string",
						Description: "New name of the label",
						Required:    true,
					},
					"messageListVisibility": {
						Value:       "",
						Type:        "string",
						Description: "Visibility of the message with this label",
						Required:    true,
					},
					"LabelListVisibility": {
						Value:       "",
						Type:        "string",
						Description: "Visibility of the label",
						Required:    true,
					},
					"type": {
						Value:       "",
						Type:        "string",
						Description: "Type of the label (User/System)",
						Required:    true,
					},
				},
			},
			{
				Name:    "Delete a label",
				RefName: "gmail/deleteLabel",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the label",
						Required:    true,
					},
				},
			},
			{
				Name:    "Create a comment in a file",
				RefName: "gmail/createComment",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the file where to put the comment",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "Content of the comment",
						Required:    true,
					},
				},
			},
			{
				Name:    "Delete a comment in a file",
				RefName: "gmail/deleteComment",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the file where to put the comment",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "Content of the comment",
						Required:    true,
					},
				},
			},
			{
				Name:    "Update a comment in a file",
				RefName: "gmail/updateComment",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the file where to put the comment",
						Required:    true,
					},
					"old_content": {
						Value:       "",
						Type:        "string",
						Description: "Old content of the comment",
						Required:    true,
					},
					"new_content": {
						Value:       "",
						Type:        "string",
						Description: "New content of the comment",
						Required:    true,
					},
				},
			},
			{
				Name:    "Create an empty file",
				RefName: "gmail/createEmptyFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the new file",
						Required:    true,
					},
					"description": {
						Value:       "",
						Type:        "string",
						Description: "Description of the file created",
						Required:    true,
					},
				},
			},
			{
				Name:    "Delete a file",
				RefName: "gmail/deleteFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the file",
						Required:    true,
					},
				},
			},
			{
				Name:    "Update a file",
				RefName: "gmail/updateFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"old_file_name": {
						Value:       "",
						Type:        "string",
						Description: "Old name of the file",
						Required:    true,
					},
					"new_file_name": {
						Value:       "",
						Type:        "string",
						Description: "New name of the file",
						Required:    true,
					},
					"description": {
						Value:       "",
						Type:        "string",
						Description: "Description of the file",
						Required:    true,
					},
				},
			},
			{
				Name:    "Copy a file",
				RefName: "gmail/copyFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"src_file_name": {
						Value:       "",
						Type:        "string",
						Description: "Source name of the file",
						Required:    true,
					},
					"dest_file_name": {
						Value:       "",
						Type:        "string",
						Description: "Destination name of the file",
						Required:    true,
					},
					"description": {
						Value:       "",
						Type:        "string",
						Description: "Description of the file",
						Required:    true,
					},
				},
			},
			{
				Name:    "Watch email received by an user",
				RefName: "watchme",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{},
			},
			{
				Name:    "Watch a specific file",
				RefName: "watchFile",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the file",
						Required:    true,
					},
				},
			},
			{
				Name:    "Watch Changes for all file ressources",
				RefName: "watchChanges",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{},
			},
		},
	}
	return status, nil
}
