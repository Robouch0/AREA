//
// EPITECH PROJECT, 2025
// AREA
// File description:
// discordClientList
//

package discord_client

import (
	IServ "area/gRPC/api/serviceInterface"
)

func (disCli *DiscordClient) ListServiceStatus() (*IServ.ServiceStatus, error) {
	status := &IServ.ServiceStatus{
		Name:    "Discord",
		RefName: "discord",

		Microservices: []IServ.MicroserviceDescriptor{
			{
				Name:    "Create a message in a channel",
				RefName: "createMsg",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"channel": {
						Value:       "",
						Type:        "string",
						Description: "Channel Discord where to put this message",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "Content of the new message",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO Matthieu
			},
			{
				Name:    "Edit a message",
				RefName: "editMsg",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"channel": {
						Value:       "",
						Type:        "string",
						Description: "Channel Discord where to edit the message",
						Required:    true,
					},
					"message_id": {
						Value:       "",
						Type:        "string",
						Description: "Message Identifier (available in discord app)",
						Required:    true,
					},
					"content": {
						Value:       "",
						Type:        "string",
						Description: "Content of the message",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO Matthieu
			},
			{
				Name:    "Delete a message",
				RefName: "deleteMsg",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"channel": {
						Value:       "",
						Type:        "string",
						Description: "Channel Discord where to delete the message",
						Required:    true,
					},
					"message_id": {
						Value:       "",
						Type:        "string",
						Description: "Message Identifier (available in discord app)",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO Matthieu
			},
			{
				Name:    "Create a reaction on a message",
				RefName: "createReact",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"channel": {
						Value:       "",
						Type:        "string",
						Description: "Channel Discord where to create a reaction",
						Required:    true,
					},
					"message_id": {
						Value:       "",
						Type:        "string",
						Description: "Message Identifier (available in discord app)",
						Required:    true,
					},
					"emoji": {
						Value:       "",
						Type:        "string",
						Description: "Emoji to send",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO Matthieu
			},
			{
				Name:    "Delete all reactions on a message",
				RefName: "deleteAllreacts",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"channel": {
						Value:       "",
						Type:        "string",
						Description: "Channel Discord where to delete the reactions",
						Required:    true,
					},
					"message_id": {
						Value:       "",
						Type:        "string",
						Description: "Message Identifier (available in discord app)",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO Matthieu
			},
			{
				Name:    "Delete selected reactions on a message",
				RefName: "deleteReact",
				Type:    "reaction",

				Ingredients: map[string]IServ.IngredientDescriptor{
					"channel": {
						Value:       "",
						Type:        "string",
						Description: "Channel Discord where to delete the reactions",
						Required:    true,
					},
					"message_id": {
						Value:       "",
						Type:        "string",
						Description: "Message Identifier (available in discord app)",
						Required:    true,
					},
					"emoji": {
						Value:       "",
						Type:        "string",
						Description: "Emoji to send",
						Required:    true,
					},
				},
				PipelineAvailable: []string{}, // TODO Matthieu
			},
		},
	}
	return status, nil
}
