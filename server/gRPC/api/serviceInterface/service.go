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

type ClientService interface {
	SendAction(body map[string]any, actionId int) (*ActionResponseStatus, error)
	// TriggerReaction(action int, prevRes string) (int, error)
}
