//
// EPITECH PROJECT, 2024
// AREA
// File description:
// service
//

package serviceinterface

import (
	"area/models"
	"net/http"
)

// Datas sent by a remote service to the app (as a webhook callback)
type WebhookInfos struct {
	Payload map[string]any `json:"payload,omitempty"`
	Header  http.Header    `json:"header,omitempty"`
}

// Alias for maps of reactions functions
type ReactionLauncher = map[string]func(ingredients map[string]any, prevOutput []byte, userID int) (*ReactionResponseStatus, error)

// Alias for maps of actions functions
type ActionLauncher = map[string]func(scenario models.AreaScenario, actionId, userID int) (*ActionResponseStatus, error)

type ClientService interface {
	// List the status of the current service, including the microservice currently supported
	ListServiceStatus() (*ServiceStatus, error)

	// Send an action that a service should watch
	SendAction(body models.AreaScenario, actionId, userID int) (*ActionResponseStatus, error)

	// Deactivate an area
	//
	// Parameter id must be used to identify the area, and userID identifies the user
	DeactivateArea(id, userID int) (*DeactivateResponseStatus, error)

	// Trigger a specific reaction of an user
	TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*ReactionResponseStatus, error)

	// Trigger the webhook callback sent by a remote service
	TriggerWebhook(webhook *WebhookInfos, microservice string, action_id int) (*WebHookResponseStatus, error)
}
