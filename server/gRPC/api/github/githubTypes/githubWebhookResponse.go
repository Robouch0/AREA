//
// EPITECH PROJECT, 2024
// AREA [WSL: Ubuntu]
// File description:
// githubWebhookResponse
//

package githubtypes

type GithubEvents struct {
	Hook GithubWebHookResponse `json:"hook"`
}

type GithubWebHookResponse struct {
	Events []string `json:"events"`
	Id     int64    `json:"id"`
}
