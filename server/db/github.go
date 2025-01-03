//
// EPITECH PROJECT, 2024
// AREA [WSL: Ubuntu]
// File description:
// github
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type GithubDB struct {
	Db *bun.DB
}

func InitGithubDb() (*GithubDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Github)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &GithubDB{Db: db}, err
}

// If we need to do activation/disactivation do like dateTimeDB

func (hf *GithubDB) StoreNewGithub(newData *models.Github) (*models.Github, error) {
	_, err := hf.Db.NewInsert().
		Model(newData).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (hf *GithubDB) GetGithubByActionID(actionID uint) (*models.Github, error) {
	allDatas := new(models.Github)
	err := hf.Db.NewSelect().
		Model(allDatas).
		Where("action_id = ?", actionID).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

// Activate or desactivate an action based on actionID and the boolean activated
func (git *GithubDB) SetActivateByActionID(activated bool, userID, actionID uint) (*models.Github, error) {
	return SetActivateByActionID[models.Github](git.Db, activated, userID, actionID)
}

func (hf *GithubDB) GetAllActionsActivated() (*[]models.Github, error) {
	allDatas := new([]models.Github)
	err := hf.Db.NewSelect().
		Model(allDatas).
		Where("activated = TRUE").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (hf *GithubDB) GetAllDTActions() (*[]models.Github, error) {
	return GetAll[models.Github](hf.Db)
}
