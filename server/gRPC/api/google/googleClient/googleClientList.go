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
				PipelineAvailable: []string{"to", "subject", "body_message"},
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
				PipelineAvailable: []string{},
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
				PipelineAvailable: []string{"subject"},
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
				PipelineAvailable: []string{"subject"},
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
				PipelineAvailable: []string{"name", "messageListVisibility", "LabelListVisibility", "type"},
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
				PipelineAvailable: []string{"old_name", "new_name", "messageListVisibility", "LabelListVisibility", "type"},
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
				PipelineAvailable: []string{"name"},
			},
			{
				Name:    "Create a comment in a file",
				RefName: "drive/createComment",
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
				PipelineAvailable: []string{"file_name", "content"},
			},
			{
				Name:    "Delete a comment in a file",
				RefName: "drive/deleteComment",
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
				PipelineAvailable: []string{"file_name", "content"},
			},
			{
				Name:    "Update a comment in a file",
				RefName: "drive/updateComment",
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
				PipelineAvailable: []string{"file_name", "old_content", "new_content"},
			},
			{
				Name:    "Create an empty file",
				RefName: "drive/createEmptyFile",
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
				PipelineAvailable: []string{"file_name", "description"},
			},
			{
				Name:    "Delete a file",
				RefName: "drive/deleteFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_name": {
						Value:       "",
						Type:        "string",
						Description: "Name of the file",
						Required:    true,
					},
				},
				PipelineAvailable: []string{"file_name"},
			},
			{
				Name:    "Update a file",
				RefName: "drive/updateFile",
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
				PipelineAvailable: []string{"old_file_name", "new_file_name", "description"},
			},
			{
				Name:    "Copy a file",
				RefName: "drive/copyFile",
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
				PipelineAvailable: []string{"src_file_name", "dest_file_name", "description"},
			},
			{
				Name:    "Watch email received by an user",
				RefName: "watchme", // No subdomain because of google cloud inner functionning
				Type:    "action",

				Ingredients:       map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{"data", "messageId", "publishTime"},
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
				PipelineAvailable: []string{},
			},
			{
				Name:    "Watch Changes for all file ressources",
				RefName: "watchChanges",
				Type:    "action",

				Ingredients:       map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{},
			},
		},
	}
	return status, nil
}
