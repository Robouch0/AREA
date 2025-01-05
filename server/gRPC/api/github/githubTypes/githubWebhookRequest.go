//
// EPITECH PROJECT, 2024
// AREA [WSL: Ubuntu]
// File description:
// githubWebhookRequest
//

package githubtypes

type GithubConfig struct {
	Url     string `json:"url"`
	Content string `json:"content_type"`
}

type GitWebHookRequest struct {
	Event  []string     `json:"events"`
	Active bool         `json:"active"`
	Config GithubConfig `json:"config"`
}
