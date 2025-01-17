//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlabWebhookResponse
//

package gitlabtypes

type GitabProject struct {
	ProjectId int `json:"id,omitempty"`
}

type ObjectAttribute struct {
	Action string `json:"action,omitempty"`
}

type GitlabWebHookResponse struct {
	EventName  string          `json:"event_name,omitempty"`
	EventType  string          `json:"event_type,omitempty"`
	ObjectKind string          `json:"object_kind,omitempty"`
	Project    GitabProject    `json:"project,omitempty"`
	Object     ObjectAttribute `json:"object_attributes,omitempty"`
	Action     string          `json:"action,omitempty"`
}
