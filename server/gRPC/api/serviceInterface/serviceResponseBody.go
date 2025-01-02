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
}

type WebHookResponseStatus struct {
	Description string `json:"description"`
}

type DeactivateResponseStatus struct {
}
