//
// EPITECH PROJECT, 2024
// AREA
// File description:
// gitlabClientList
//

package gitlab_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (git *GitlabClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Gitlab",
		RefName: "gitlab",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Create a file on a repository",
				RefName: "createFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_path": {
						Value:       "",
						Type:        "string",
						Description: "URL-encoded full path to new file",
						Required:    true,
					},
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID or URL-encoded path of the project",
						Required:    true,
					},
					"branch": {
						Value:       "",
						Type:        "string",
						Description: "Name of the branch which holds the new file",
						Required:    true,
					},
					"commit_message": {
						Value:       "",
						Type:        "string",
						Description: "Message of the commit",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "File's content",
						Required:    true,
					},
				},
			},
			{
				Name:    "Update a file on a repository",
				RefName: "updateFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_path": {
						Value:       "",
						Type:        "string",
						Description: "URL-encoded full path to new file",
						Required:    true,
					},
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID or URL-encoded path of the project",
						Required:    true,
					},
					"branch": {
						Value:       "",
						Type:        "string",
						Description: "Name of the branch which holds the new file",
						Required:    true,
					},
					"commit_message": {
						Value:       "",
						Type:        "string",
						Description: "Message of the commit",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "File's content",
						Required:    true,
					},
				},
			},
			{
				Name:    "Delete a file on a repository",
				RefName: "deleteFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"file_path": {
						Value:       "",
						Type:        "string",
						Description: "URL-encoded full path to new file",
						Required:    true,
					},
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID or URL-encoded path of the project",
						Required:    true,
					},
					"branch": {
						Value:       "",
						Type:        "string",
						Description: "Name of the branch which holds the new file",
						Required:    true,
					},
					"commit_message": {
						Value:       "",
						Type:        "string",
						Description: "Message of the commit",
						Required:    true,
					},
				},
			},
			{
				Name:    "Mark an item of a todo list as done",
				RefName: "markItemDone",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID of to-do item",
						Required:    true,
					},
				},
			},
			{
				Name:    "Mark all Items of your todo list as done",
				RefName: "markAllItemDone",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{},
			},
		},
	}
	return status, nil
}
