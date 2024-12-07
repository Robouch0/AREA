//
// EPITECH PROJECT, 2024
// AREA
// File description:
// service
//

package serviceinterface

type ActionResponseStatus struct {
	Description string `json:"description"`
	ActionID    int    `json:"action_id"`
}

type ReactionResponseStatus struct {
	Description string `json:"description"`
	ReactionID  int    `json:"reaction_id"`
}

type ClientService interface {
	SendAction(body map[string]any, actionId int) (*ActionResponseStatus, error)
	// prevOutput is an array of byte because output can be raw
	TriggerReaction(ingredients map[string]any, microservice string, prevOutput []byte) (*ReactionResponseStatus, error)
}
