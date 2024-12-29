//
// EPITECH PROJECT, 2024
// AREA
// File description:
// area
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type AreaDB struct {
	Db *bun.DB
}

func InitAreaDb() (*AreaDB, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Area)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &AreaDB{Db: db}, err
}

func GetAreaDb() *AreaDB {
	db := initDB()
	return &AreaDB{Db: db}
}

func (area *AreaDB) SubmitNewArea(newArea *models.Area) (*models.Area, error) {
	_, err := area.Db.NewInsert().
		Model(newArea).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newArea, nil
}

func (area *AreaDB) InsertNewArea(UserID uint, OneShot bool) (*models.Area, error) {
	newArea := &models.Area{
		UserID:  UserID,
		OneShot: OneShot,
	}
	_, err := area.Db.NewInsert().
		Model(newArea).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newArea, nil
}

func (area *AreaDB) GetFullAreaByUserID(userID uint) (*[]models.Area, error) {
	us := new([]models.Area)
	err := area.Db.NewSelect().
		Model(us).
		Where("user_id = ?", userID).
		Relation("Action", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("user_id = ?", userID)
		}).
		Relation("Reactions", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("user_id = ?", userID)
		}).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return us, nil
}

func (area *AreaDB) GetAreaByID(AreaID uint) (*models.Area, error) {
	return GetByID[models.Area](area.Db, AreaID)
}

func (area *AreaDB) GetAreaByActionID(ActionID uint) (*models.Area, error) {
	data := new(models.Area)
	err := area.Db.NewSelect().
		Model(data).
		Relation("Action", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("area_id = ?", ActionID)
		}).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (area *AreaDB) GetUserAreaByActionID(UserId, ActionID uint) (*models.Area, error) {
	data := new(models.Area)
	err := area.Db.NewSelect().
		Model(data).
		Relation("Action", func(sq *bun.SelectQuery) *bun.SelectQuery {
			return sq.Where("area_id = ?", ActionID)
		}).
		Where("user_id = ?", UserId).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return data, nil
}
