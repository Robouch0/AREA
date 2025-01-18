//
// EPITECH PROJECT, 2024
// AREA
// File description:
// actions
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type ActionsDb struct {
	Db *bun.DB
}

func InitActionsDb() (*ActionsDb, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Actions)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &ActionsDb{Db: db}, err
}

func GetActionDb() *ActionsDb {
	db := initDB()
	return &ActionsDb{Db: db}
}

func (action *ActionsDb) SubmitNewAction(newAction *models.Actions) (*models.Actions, error) {
	_, err := action.Db.NewInsert().
		Model(newAction).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newAction, nil
}

func (action *ActionsDb) InsertNewAction(ActionInfo *models.Action, AreaID uint) (*models.Actions, error) {
	newAction := &models.Actions{
		AreaID: AreaID,
		Action: ActionInfo,
	}
	_, err := action.Db.NewInsert().
		Model(newAction).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newAction, nil
}

func (action *ActionsDb) DeleteByActionID(actionID uint) error {
	data := new(models.Actions)
	_, err := action.Db.NewDelete().
		Model(data).
		Where("id = ?", actionID).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (action *ActionsDb) GetActionByID(ID uint) (*models.Actions, error) {
	return GetByID[models.Actions](action.Db, ID)
}
