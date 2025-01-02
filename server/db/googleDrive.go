//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleDrive
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type GoogleDriveDB struct {
	Db *bun.DB
}

func InitGoogleDriveDB() (*GoogleDriveDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Drive)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &GoogleDriveDB{Db: db}, err
}

// If we need to do activation/disactivation do like dateTimeDB

func (google *GoogleDriveDB) StoreNewGWatch(newData *models.Drive) (*models.Drive, error) {
	_, err := google.Db.NewInsert().
		Model(newData).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (google *GoogleDriveDB) GetActionByID(action_id string) (*models.Drive, error) {
	allDatas := new(models.Drive)
	err := google.Db.NewSelect().
		Model(allDatas).
		Where("action_id = ?", action_id).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (google *GoogleDriveDB) GetByChannelID(channelID string) (*models.Drive, error) {
	allDatas := new(models.Drive)
	err := google.Db.NewSelect().
		Model(allDatas).
		Where("channel_id = ?", channelID).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (google *GoogleDriveDB) GetAllActionsActivated() (*[]models.Drive, error) {
	allDatas := new([]models.Drive)
	err := google.Db.NewSelect().
		Model(allDatas).
		Where("activated = TRUE").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (google *GoogleDriveDB) GetAllActions() (*[]models.Drive, error) {
	return GetAll[models.Drive](google.Db)
}
