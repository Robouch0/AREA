//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFace
//

package models

import "github.com/uptrace/bun"

type HuggingFace struct {
	bun.BaseModel `bun:"table:dateTime,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID  uint `bun:"action_id" json:"action_id"`
	UserID    uint `bun:"user_id" json:"user_id"`
	Activated bool `bun:"activated" json:"activated,omitempty"`

	RepoType string `bun:"repo_type" json:"type"` /* Type of repository (model, dataset, etc.) */
	RepoName string `bun:"repo_name" json:"name"` /* Name of the repository as it is named in HF */

	RepoAction string `bun:"repo_action" json:"action"` /* Action to track in the event json on the webhook payload */
	RepoScope  string `bun:"repo_scope" json:"scope"`   /* Scope of the webhook sent in the payload */

	// Action checks pull request only if scope is discussion
	IsPullRequest bool `bun:"is_pull_request" json:"is_pull_request,omitempty"`
}
