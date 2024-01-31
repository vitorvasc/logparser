package port

import "logparser/internal/core/domain"

type MatchRepository interface {
	SaveMatch(match *domain.Match) error
	FindMatchByID(id string) (*domain.Match, error)
}

type CreateMatchService interface {
	BulkCreate(historyList []*domain.MatchHistory) []*domain.BulkCreationResult
	Create(history *domain.MatchHistory) (*domain.Match, error)
}
