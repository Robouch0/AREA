//
// EPITECH PROJECT, 2024
// AREA
// File description:
// dateTime
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type DateTimeDB struct {
	Db *bun.DB
}

func InitDateTimeDb() (*DateTimeDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.DateTime)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &DateTimeDB{Db: db}, err
}

// Upsert a new DateTime action (insert if it does not exist else it updates it)
func (dt *DateTimeDB) InsertNewDTAction(dtAction *models.DateTime) (*models.DateTime, error) {
	_, err := dt.Db.NewInsert().
		Model(dtAction).
		On("CONFLICT (id) DO UPDATE").
		Set("activated = TRUE").
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return dtAction, nil
}

// Activate or desactivate an action based on actionID and the boolean activated
func (dt *DateTimeDB) SetActivateByActionID(activated bool, userID, actionID uint) (*models.DateTime, error) {
	return SetActivateByActionID[models.DateTime](dt.Db, activated, userID, actionID)
}

func (dt *DateTimeDB) GetByActionID(actionID uint) (*models.DateTime, error) {
	return GetByActionID[models.DateTime](dt.Db, actionID)
}

func (dt *DateTimeDB) GetAllDTActionsActivated() (*[]models.DateTime, error) {
	allDatas := new([]models.DateTime)
	err := dt.Db.NewSelect().
		Model(allDatas).
		Where("activated = TRUE").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (dt *DateTimeDB) GetAllDTActions() (*[]models.DateTime, error) {
	return GetAll[models.DateTime](dt.Db)
}
