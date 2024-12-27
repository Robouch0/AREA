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

// Do enable/disable
// GetAll
// ...
