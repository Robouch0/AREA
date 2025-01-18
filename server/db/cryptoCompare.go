//
// EPITECH PROJECT, 2024
// AREA
// File description:
// weatherCondition
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type CryptoCompareDB struct {
	Db *bun.DB
}

func InitCryptoCompareDb() (*CryptoCompareDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.CryptoCompare)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &CryptoCompareDB{Db: db}, err
}

func (crypto *CryptoCompareDB) InsertNewCryptoCompare(data *models.CryptoCompare) (*models.CryptoCompare, error) {
	_, err := crypto.Db.NewInsert().
		Model(data).
		On("CONFLICT (id) DO UPDATE").
		Set("activated = TRUE").
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (crypto *CryptoCompareDB) GetActionsByType(actionType models.CryptoActionType) (*[]models.CryptoCompare, error) {
	allDatas := new([]models.CryptoCompare)
	err := crypto.Db.NewSelect().
		Model(allDatas).
		Where("activated = TRUE").
		Where("action_type = ?", actionType).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allDatas, nil
}

// Activate or desactivate an action based on actionID and the boolean activated
func (crypto *CryptoCompareDB) SetActivateByActionID(activated bool, userID, actionID uint) (*models.CryptoCompare, error) {
	return SetActivateByActionID[models.CryptoCompare](crypto.Db, activated, userID, actionID)
}

func (crypto *CryptoCompareDB) DeleteByActionID(userID, actionID uint) error {
	return DeleteUserActionByActionID[models.CryptoCompare](crypto.Db, userID, actionID)
}
