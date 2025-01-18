//
// EPITECH PROJECT, 2024
// AREA
// File description:
// huggingFace
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type HuggingFaceDB struct {
	Db *bun.DB
}

func InitHuggingFaceDb() (*HuggingFaceDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.HuggingFace)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &HuggingFaceDB{Db: db}, err
}

// If we need to do activation/disactivation do like dateTimeDB

func (hf *HuggingFaceDB) StoreNewHF(newData *models.HuggingFace) (*models.HuggingFace, error) {
	_, err := hf.Db.NewInsert().
		Model(newData).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (hf *HuggingFaceDB) GetHfByActionID(actionID uint) (*models.HuggingFace, error) {
	allDatas := new(models.HuggingFace)
	err := hf.Db.NewSelect().
		Model(allDatas).
		Where("action_id = ?", actionID).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (hf *HuggingFaceDB) GetAllActionsActivated() (*[]models.HuggingFace, error) {
	allDatas := new([]models.HuggingFace)
	err := hf.Db.NewSelect().
		Model(allDatas).
		Where("activated = TRUE").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (hf *HuggingFaceDB) GetAllDTActions() (*[]models.HuggingFace, error) {
	return GetAll[models.HuggingFace](hf.Db)
}

func (hf *HuggingFaceDB) DeleteByActionID(userID, actionID uint) error {
	return DeleteUserActionByActionID[models.HuggingFace](hf.Db, userID, actionID)
}
