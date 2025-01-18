//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// gitlab
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type GitlabDB struct {
	Db *bun.DB
}

func InitGitlabDb() (*GitlabDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Gitlab)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &GitlabDB{Db: db}, err
}

// If we need to do activation/disactivation do like dateTimeDB

func (git *GitlabDB) StoreNewGithub(newData *models.Gitlab) (*models.Gitlab, error) {
	_, err := git.Db.NewInsert().
		Model(newData).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (git *GitlabDB) GetGithubByActionID(actionID uint) (*models.Gitlab, error) {
	allDatas := new(models.Gitlab)
	err := git.Db.NewSelect().
		Model(allDatas).
		Where("action_id = ?", actionID).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

// Activate or desactivate an action based on actionID and the boolean activated
func (git *GitlabDB) SetActivateByActionID(activated bool, userID, actionID uint) (*models.Gitlab, error) {
	return SetActivateByActionID[models.Gitlab](git.Db, activated, userID, actionID)
}
func (git *GitlabDB) GetAllActionsActivated() (*[]models.Gitlab, error) {
	allDatas := new([]models.Gitlab)
	err := git.Db.NewSelect().
		Model(allDatas).
		Where("activated = TRUE").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (git *GitlabDB) GetAllDTActions() (*[]models.Gitlab, error) {
	return GetAll[models.Gitlab](git.Db)
}

func (git *GitlabDB) DeleteByActionID(userID, actionID uint) error {
	return DeleteUserActionByActionID[models.Gitlab](git.Db, userID, actionID)
}
