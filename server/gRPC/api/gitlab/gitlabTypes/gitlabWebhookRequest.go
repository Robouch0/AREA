//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlabWebhookRequest
//

package gitlabtypes

type GitlabWebHookRequest struct {
	Url          string `json:"url"`
	PushEvent    bool   `json:"push_events"`
	IssuesEvent  bool   `json:"issues_events"`
	TagEvent     bool   `json:"tag_push_events"`
	ReleaseEvent bool   `json:"releases_events"`
	MergeEvent   bool   `json:"merge_requests_events"`
}
