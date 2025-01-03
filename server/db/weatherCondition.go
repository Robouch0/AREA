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

type WeatherConditionDB struct {
	Db *bun.DB
}

func InitWeatherConditionDb() (*WeatherConditionDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.WeatherCondition)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &WeatherConditionDB{Db: db}, err
}

func (wDb *WeatherConditionDB) InsertNewWeatherCondition(data *models.WeatherCondition) (*models.WeatherCondition, error) {
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

func (wDb *WeatherConditionDB) GetActionsByType(actionType models.WeatherActionType) (*[]models.WeatherCondition, error) {
	allDatas := new([]models.WeatherCondition)
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

func (wDb *WeatherConditionDB) UpdateUserIsDay(actionID, IsDay int) error {
	var wData models.WeatherCondition

	_, err := wDb.Db.NewUpdate().
		Model(&wData).
		Set("is_day = ?", IsDay).
		Where("action_id = ?", actionID).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Activate or desactivate an action based on actionID and the boolean activated
func (weather *WeatherConditionDB) SetActivateByActionID(activated bool, userID, actionID uint) (*models.WeatherCondition, error) {
	return SetActivateByActionID[models.WeatherCondition](weather.Db, activated, userID, actionID)
}
