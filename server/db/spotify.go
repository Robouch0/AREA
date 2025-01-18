//
// EPITECH PROJECT, 2025
// AREA [WSL: Ubuntu]
// File description:
// spotify
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type SpotifyDB struct {
	Db *bun.DB
}

func InitSpotifyDb() (*SpotifyDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Spotify)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &SpotifyDB{Db: db}, err
}

func (wDb *SpotifyDB) InsertNewSpotify(data *models.Spotify) (*models.Spotify, error) {
	_, err := wDb.Db.NewInsert().
		Model(data).
		On("CONFLICT (id) DO UPDATE").
		Set("activated = TRUE").
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (wDb *SpotifyDB) GetActionsByType(actionType models.SpotifyActionType) (*[]models.Spotify, error) {
	allDatas := new([]models.Spotify)
	err := wDb.Db.NewSelect().
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
func (weather *SpotifyDB) SetActivateByActionID(activated bool, userID, actionID uint) (*models.Spotify, error) {
	return SetActivateByActionID[models.Spotify](weather.Db, activated, userID, actionID)
}

func (spot *SpotifyDB) DeleteByActionID(userID, actionID uint) error {
	return DeleteUserActionByActionID[models.Spotify](spot.Db, userID, actionID)
}
