//
// EPITECH PROJECT, 2024
// AREA [WSL: Ubuntu]
// File description:
// githubWebhookResponse
//

package githubtypes

type GithubWebHookResponse struct {
	Event      []string    `json:"event"`
}