//
// EPITECH PROJECT, 2025
// AREA
// File description:
// miro
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type MiroDB struct {
	Db *bun.DB
}

func InitMiroDb() (*MiroDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Miro)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &MiroDB{Db: db}, err
}

// If we need to do activation/disactivation do like dateTimeDB

func (miro *MiroDB) StoreNewMiro(newData *models.Miro) (*models.Miro, error) {
	_, err := miro.Db.NewInsert().
		Model(newData).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (miro *MiroDB) GetMiroByActionID(actionID uint) (*models.Miro, error) {
	allDatas := new(models.Miro)
	err := miro.Db.NewSelect().
		Model(allDatas).
		Where("action_id = ?", actionID).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

// Activate or desactivate an action based on actionID and the boolean activated
func (git *MiroDB) SetActivateByActionID(activated bool, userID, actionID uint) (*models.Miro, error) {
	return SetActivateByActionID[models.Miro](git.Db, activated, userID, actionID)
}

func (miro *MiroDB) GetAllActionsActivated() (*[]models.Miro, error) {
	allDatas := new([]models.Miro)
	err := miro.Db.NewSelect().
		Model(allDatas).
		Where("activated = TRUE").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (miro *MiroDB) GetAllMiroActions() (*[]models.Miro, error) {
	return GetAll[models.Miro](miro.Db)
}

func (miro *MiroDB) DeleteByActionID(actionID uint) error {
	return DeleteByActionID[models.Miro](miro.Db, actionID)
}
