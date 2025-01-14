//
// EPITECH PROJECT, 2024
// AREA [WSL: Ubuntu]
// File description:
// githubClientList
//

package github

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (git *GithubClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Github",
		RefName: "github",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Trigger every new push of a repository",
				RefName: "triggerPush",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Trigger every time an issue is created",
				RefName: "triggerIssue",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Trigger every time an issue is closed",
				RefName: "triggerIssueClose",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Trigger every time a pull request is opened",
				RefName: "triggerPr",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Trigger every time a pull request is closed",
				RefName: "triggerPrClose",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Trigger every time a branch or a tag is created",
				RefName: "triggerCreate",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Trigger every time a branch or a tag is deleted",
				RefName: "triggerDelete",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Trigger every time the repository is forked",
				RefName: "triggerFork",
				Type:    "action",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Update Repository Informations",
				RefName: "updateRepo",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
					"name": {
						Value:       "",
						Type:        "string",
						Description: "New name for the repository",
						Required:    true,
					},
					"description": {
						Value:       "",
						Type:        "string",
						Description: "New description for the repository",
						Required:    true,
					},
				},
			},
			{
				Name:    "Update a file in a repository",
				RefName: "updateFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
					"path": {
						Value:       "",
						Type:        "string",
						Description: "Path of the file in the repository",
						Required:    true,
					},
					"message": {
						Value:       "",
						Type:        "string",
						Description: "Commit message",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "New content of the file",
						Required:    true,
					},
				},
			},
			{
				Name:    "Delete Repository File",
				RefName: "deleteFile",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"owner": {
						Value:       "",
						Type:        "string",
						Description: "Owner of the repository",
						Required:    true,
					},
					"repo": {
						Value:       "",
						Type:        "string",
						Description: "Name of the repository",
						Required:    true,
					},
					"path": {
						Value:       "",
						Type:        "string",
						Description: "Path of the file in the repository",
						Required:    true,
					},
					"message": {
						Value:       "",
						Type:        "string",
						Description: "Commit message",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}
