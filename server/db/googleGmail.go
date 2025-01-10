//
// EPITECH PROJECT, 2024
// AREA
// File description:
// googleGmail
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type GoogleGmailDB struct {
	Db *bun.DB
}

func InitGoogleGmailDb() (*GoogleGmailDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Gmail)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &GoogleGmailDB{Db: db}, err
}

// If we need to do activation/disactivation do like dateTimeDB

func (google *GoogleGmailDB) StoreNewGWatch(newData *models.Gmail) (*models.Gmail, error) {
	_, err := google.Db.NewInsert().
		Model(newData).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newData, nil
}

func (google *GoogleGmailDB) GetByEmail(email string) (*models.Gmail, error) {
	allDatas := new(models.Gmail)
	err := google.Db.NewSelect().
		Model(allDatas).
		Where("email_address = ?", email).
		Where("activated = TRUE").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (google *GoogleGmailDB) SetFirstTime(action_id uint, firstTime bool) (*models.Gmail, error) {
	allDatas := new(models.Gmail)
	_, err := google.Db.NewUpdate().
		Model(allDatas).
		Set("first_time = ?", firstTime).
		Where("action_id = ?", action_id).
		Exec(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

// Activate or desactivate an action based on actionID and the boolean activated
func (google *GoogleGmailDB) SetActivateByActionID(activated bool, userID, actionID uint) (*models.Gmail, error) {
	return SetActivateByActionID[models.Gmail](google.Db, activated, userID, actionID)
}

func (google *GoogleGmailDB) GetAllActionsActivated() (*[]models.Gmail, error) {
	allDatas := new([]models.Gmail)
	err := google.Db.NewSelect().
		Model(allDatas).
		Where("activated = TRUE").
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

func (google *GoogleGmailDB) GetAllActions() (*[]models.Gmail, error) {
	return GetAll[models.Gmail](google.Db)
}
