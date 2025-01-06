//
// EPITECH PROJECT, 2025
// AREA
// File description:
// type
//

package miro_webhook

type MiroData struct {
	Content string `json:"content,omitempty"`
	Shape   string `json:"shape,omitempty"`
}

type MiroItem struct {
	ID   string   `json:"id"`
	Data MiroData `json:"data,omitempty"`
}

type MiroWebhookPayload struct {
	BoardId string   `json:"board_id"`
	Type    string   `json:"type"`
	Item    MiroItem `json:"item"`
}

type MiroChallenge struct {
	Challenge string `json:"challenge"`
}
