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
				Name:    "Triggers every push",
				RefName: "triggerPush",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID or URL-encoded path of the project",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO Matthieu
			},
			{
				Name:    "Triggers at every creation of a new issue on a project",
				RefName: "triggerIssue",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID or URL-encoded path of the project",
						Required:    true,
					},
				},
			},
			{
				Name:    "Triggers on every tag event on a project",
				RefName: "triggerTag",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID or URL-encoded path of the project",
						Required:    true,
					},
				},
			},
			{
				Name:    "Triggers on every new release on a project",
				RefName: "triggerRelease",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID or URL-encoded path of the project",
						Required:    true,
					},
				},
			},
			{
				Name:    "Trigger every merge request event on a project",
				RefName: "triggerMerge",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"id": {
						Value:       "",
						Type:        "string",
						Description: "The ID or URL-encoded path of the project",
						Required:    true,
					},
				},
			},
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
				PipelineAvailable: []string{}, // TODO Matthieu
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
				PipelineAvailable: []string{}, // TODO Matthieu
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
				PipelineAvailable: []string{}, // TODO Matthieu
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
				PipelineAvailable: []string{}, // TODO Matthieu
			},
			{
				Name:    "Mark all Items of your todo list as done",
				RefName: "markAllItemDone",
				Type:    "reaction",

				Ingredients:       map[string]IServ.IngredientDescriptor{},
				PipelineAvailable: []string{}, // TODO Matthieu
			},
		},
	}
	return status, nil
}
