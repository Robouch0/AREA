//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlab
//

package models

import "github.com/uptrace/bun"

type GlAction string

type GlType string

const (
	GlPush    GlAction = "push"
	GlTag              = "tag_push"
	GlRelease          = "release"
	GlIssue            = "issue"
	GlMergeR           = "merge_request"
)

const (
	Glopen   GlType = "open"
	Glcreate 		= "create"
	GlEmpty         = ""
)

type Gitlab struct {
	bun.BaseModel `bun:"table:github,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID  uint `bun:"action_id" json:"action_id"`
	UserID    uint `bun:"user_id" json:"user_id"`
	Activated bool `bun:"activated" json:"activated,omitempty"`

	RepoID     string   `bun:"repo_owner" json:"repo_id"` /* Name of the owner of the repository */
	RepoAction GlAction `bun:"repo_action" json:"action"` /* Action to track in the event json on the webhook payload */
	ActionType GlType   `bun:"action_type" json:"action_type"`
}
