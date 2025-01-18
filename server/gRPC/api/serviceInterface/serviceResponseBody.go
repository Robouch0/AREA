//
// EPITECH PROJECT, 2025
// AREA
// File description:
// serviceResponseBody
//

package serviceinterface

type ActionResponseStatus struct {
	Description string `json:"description"`
	ActionID    int    `json:"action_id"`
}

type ReactionResponseStatus struct {
	Description string `json:"description"`
	ReactionID  int    `json:"reaction_id"`

	Datas map[string]any
}

type WebHookResponseStatus struct {
	Description string `json:"description"`
}

type SetActivatedResponseStatus struct {
	ActionID    uint   `json:"action_id,omitempty"`
	Description string `json:"description,omitempty"`
	Activated   bool   `json:"activated,omitempty"`
}

type DeleteResponseStatus struct {
	ID uint `json:"id,omitempty"`
}
