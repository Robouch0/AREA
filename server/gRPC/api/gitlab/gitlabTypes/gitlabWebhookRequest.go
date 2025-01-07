//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlabWebhookRequest
//

package gitlabtypes

type GitlabWebHookRequest struct {
	Url     string `json:"url"`
	PushEvent bool `json:"push_events"`
}
