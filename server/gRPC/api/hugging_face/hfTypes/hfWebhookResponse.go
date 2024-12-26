//
// EPITECH PROJECT, 2024
// AREA
// File description:
// hfWebhookResponse
//

package hfType

type EventType struct {
	Action string `json:"action"`
	Scope  string `json:"scope"`
}

type HFDiscussion struct {
	ID            string `json:"id,omitempty"`
	Title         string `json:"title,omitempty"`
	IsPullRequest bool   `json:"isPullRequest,omitempty"`
}

type HFWebHookResponse struct {
	Event      EventType    `json:"event"`
	Repo       HFRepo       `json:"repo,omitempty"`
	Discussion HFDiscussion `json:"discussion,omitempty"`
}
