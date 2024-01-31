package memorydb

import (
	"errors"

	"logparser/internal/core/domain"
	"logparser/internal/core/port"
)

type repository struct {
	instance map[string]*domain.Match
}

func NewMatchRepository() port.MatchRepository {
	return &repository{
		instance: make(map[string]*domain.Match),
	}
}

func (db *repository) SaveMatch(match *domain.Match) error {
	db.instance[match.ID] = match
	return nil
}

func (db *repository) FindMatchByID(id string) (*domain.Match, error) {
	match, ok := db.instance[id]
	if !ok {
		return nil, errors.New("match not found")
	}
	return match, nil
}
