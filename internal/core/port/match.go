package port

import (
	"logparser/internal/core/domain"
	"logparser/internal/core/errors"
)

type MatchRepository interface {
	SaveMatch(match *domain.Match) errors.BaseError
	FindMatchByID(id string) (*domain.Match, errors.BaseError)
	FindAllMatches() ([]*domain.Match, errors.BaseError)
}

type CreateMatchService interface {
	BulkCreate(historyList []*domain.MatchHistory) []*domain.BulkCreationResult
	Create(history *domain.MatchHistory) (*domain.Match, error)
}

type GetMatchService interface {
	GetMatchByID(id string) (*domain.Match, errors.BaseError)
	GetAllMatches() ([]*domain.Match, errors.BaseError)
}
