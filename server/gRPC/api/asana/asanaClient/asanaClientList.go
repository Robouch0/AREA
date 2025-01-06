//
// EPITECH PROJECT, 2024
// AREA
// File description:
// trelloClientList
//

package asana_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (asana *AsanaClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Asana",
		RefName: "asana",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Create a Project on Asana",
				RefName: "createProject",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"projectName": {
						Value:       "",
						Type:        "string",
						Description: "The name of the project you want to create",
						Required:    true,
					},
					"color": {
						Value:       "",
						Type:        "string",
						Description: "The background color theme of your project",
						Required:    true,
					},
					"defaultView": {
						Value:       "",
						Type:        "string",
						Description: "The default view apply to your project (ex: board)",
						Required:    true,
					},
					"workspaceName": {
						Value:       "",
						Type:        "string",
						Description: "The targeted workspace name",
						Required:    true,
					},
				},
			},
			{
				Name:    "Create a Section on an Asana Project",
				RefName: "createSection",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"projectName": {
						Value:       "",
						Type:        "string",
						Description: "The name of the project where you want to create a section",
						Required:    true,
					},
					"sectionName": {
						Value:       "",
						Type:        "string",
						Description: "The name of the section you want to create",
						Required:    true,
					},
				},
			},
			{
				Name:    "Create a Task on an Asana Project",
				RefName: "createTask",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"projectName": {
						Value:       "",
						Type:        "string",
						Description: "The name of the project where you want to create a task",
						Required:    true,
					},
					"taskName": {
						Value:       "",
						Type:        "string",
						Description: "The name of the task you want to create",
						Required:    true,
					},
					"taskDescription": {
						Value:       "",
						Type:        "string",
						Description: "A short description of the task you want to create",
						Required:    true,
					},
					"completion": {
						Value:       "false",
						Type:        "bool",
						Description: "The completion status of the task your creating",
						Required:    true,
					},
					"dueOn": {
						Value:       0,
						Type:        "date",
						Description: "The due date of the created task",
						Required:    true,
					},
				},
			},
		},
	}
	return status, nil
}
