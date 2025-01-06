//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlab
//

package models

import "github.com/uptrace/bun"

type Gitlab struct {
	bun.BaseModel `bun:"table:github,alias:cs"`
	ID            uint `bun:"id,pk,autoincrement" json:"id,pk,autoincrement"`

	ActionID  uint `bun:"action_id" json:"action_id"`
	UserID    uint `bun:"user_id" json:"user_id"`
	Activated bool `bun:"activated" json:"activated,omitempty"`

	RepoId string `bun:"repo_owner" json:"owner"` /* Name of the owner of the repository */

	RepoAction string `bun:"repo_action" json:"action"` /* Action to track in the event json on the webhook payload */
}
