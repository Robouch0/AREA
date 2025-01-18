//
// EPITECH PROJECT, 2024
// AREA
// File description:
// reactions
//

package db

import (
	"area/models"
	"context"

	"github.com/uptrace/bun"
)

type ReactionDb struct {
	Db *bun.DB
}

func InitReactionsDb() (*ReactionDb, error) {
	db := initDB()

	_, err := db.NewCreateTable().
		Model((*models.Reactions)(nil)).
		IfNotExists().
		Exec(context.Background())

	return &ReactionDb{Db: db}, err
}

func GetReactionDb() *ReactionDb {
	db := initDB()
	return &ReactionDb{Db: db}
}

func (reaction *ReactionDb) SubmitNewReactions(newReaction []*models.Reactions) ([]*models.Reactions, error) {
	_, err := reaction.Db.NewInsert().
		Model(newReaction).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newReaction, nil
}

func (reaction *ReactionDb) SubmitNewReaction(newReaction *models.Reactions) (*models.Reactions, error) {
	_, err := reaction.Db.NewInsert().
		Model(newReaction).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newReaction, nil
}

func (reaction *ReactionDb) InsertNewReaction(newReact *models.Reaction, AreaID uint) (*models.Reactions, error) {
	newReaction := &models.Reactions{
		AreaID:   AreaID,
		Reaction: newReact,
	}
	_, err := reaction.Db.NewInsert().
		Model(newReaction).
		Exec(context.Background())
	if err != nil {
		return nil, err
	}
	return newReaction, nil
}

func (reaction *ReactionDb) GetReactionByID(ID uint) (*models.Reactions, error) {
	return GetByID[models.Reactions](reaction.Db, ID)
}

func (reaction *ReactionDb) GetReactionsByAreaID(AreaID uint) (*[]models.Reactions, error) {
	allReactions := new([]models.Reactions)
	err := reaction.Db.NewSelect().
		Model(allReactions).
		Where("area_id = ?", AreaID).
		Scan(context.Background())

	if err != nil {
		return nil, err
	}
	return allReactions, nil
}

func (reaction *ReactionDb) DeleteByAreaID(areaID uint) error {
	data := new(models.Reactions)
	_, err := reaction.Db.NewDelete().
		Model(data).
		Where("area_id = ?", areaID).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
