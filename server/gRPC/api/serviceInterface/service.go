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

type WebhookInfos struct {
	Payload map[string]any `json:"payload,omitempty"`
	Header  http.Header    `json:"header,omitempty"`
}

type MicroserviceLauncher = map[string]func(ingredients map[string]any, prevOutput []byte, userID int) (*ReactionResponseStatus, error)
type ActionLauncher = map[string]func(scenario models.AreaScenario, actionId, userID int) (*ActionResponseStatus, error)

type ClientService interface {
	ListServiceStatus() (*ServiceStatus, error)

	SendAction(body models.AreaScenario, actionId, userID int) (*ActionResponseStatus, error)

	// prevOutput is an array of byte because output can be raw
	TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*ReactionResponseStatus, error)

	TriggerWebhook(webhook *WebhookInfos, microservice string, action_id int) (*WebHookResponseStatus, error)
}
