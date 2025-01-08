//
// EPITECH PROJECT, 2024
// AREA [WSL: Ubuntu]
// File description:
// github
//

package models

import "github.com/uptrace/bun"

type GAction string

type GType string

const (
	GPush   GAction = "push"
	GDelete         = "delete"
	GFork           = "fork"
	GCreate         = "create"
	GIssue			= "issues"
)

const (
	Gopen GType = "opened"
	GClosed 	= "closed"
	GEmpty 		= ""
)

type Github struct {
	bun.BaseModel `bun:"table:github,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID  uint `bun:"action_id" json:"action_id"`
	UserID    uint `bun:"user_id" json:"user_id"`
	Activated bool `bun:"activated" json:"activated,omitempty"`

	RepoOwner string `bun:"repo_owner" json:"owner"` /* Name of the owner of the repository */
	RepoName  string `bun:"repo_name" json:"name"`   /* Name of the repository as it is named in github */

	RepoAction GAction `bun:"repo_action" json:"action"` /* Action to track in the event json on the webhook payload */
	ActionType GType   `bun:"action_type" json:"type"`
}
