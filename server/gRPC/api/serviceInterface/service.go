//
// EPITECH PROJECT, 2024
// AREA
// File description:
// service
//

package serviceinterface

import "area/models"

type ActionResponseStatus struct {
	Description string `json:"description"`
	ActionID    int    `json:"action_id"`
}

type ReactionResponseStatus struct {
	Description string `json:"description"`
	ReactionID  int    `json:"reaction_id"`
}

type WebHookResponseStatus struct {
	Description string `json:"description"`
}

type MicroserviceLauncher = map[string]func(ingredients map[string]any, prevOutput []byte, userID int) (*ReactionResponseStatus, error)
type ActionLauncher = map[string]func(scenario models.AreaScenario, actionId, userID int) (*ActionResponseStatus, error)

type ClientService interface {
	ListServiceStatus() (*ServiceStatus, error)

	SendAction(body models.AreaScenario, actionId, userID int) (*ActionResponseStatus, error)

	// prevOutput is an array of byte because output can be raw
	TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte, userID int) (*ReactionResponseStatus, error)

	TriggerWebhook(ingredients map[string]any, microservice string, action_id int) (*WebHookResponseStatus, error)
}
