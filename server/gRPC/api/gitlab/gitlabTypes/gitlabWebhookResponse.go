//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlabWebhookResponse
//

package gitlabtypes

type GitlabWebHookResponse struct {
	EventName string `json:"event_name,omitempty"`
	ProjectId int `json:"project_id,omitempty"`
}
