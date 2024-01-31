package memorydb

import (
	"logparser/internal/config/defines"
	"logparser/internal/core/domain"
	apperrors "logparser/internal/core/errors"
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

func (db *repository) SaveMatch(match *domain.Match) apperrors.BaseError {
	db.instance[match.ID] = match
	return nil
}

func (db *repository) FindMatchByID(id string) (*domain.Match, apperrors.BaseError) {
	match, ok := db.instance[id]
	if !ok {
		return nil, apperrors.NewError(defines.MatchNotFoundErrorCode, "match not found")
	}
	return match, nil
}

func (db *repository) FindAllMatches() ([]*domain.Match, apperrors.BaseError) {
	matchList := make([]*domain.Match, 0, len(db.instance))
	for _, match := range db.instance {
		matchList = append(matchList, match)
	}
	return matchList, nil
}
