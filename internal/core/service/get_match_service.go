package service

import (
	"logparser/internal/config/defines"
	"logparser/internal/core/domain"
	"logparser/internal/core/errors"
	"logparser/internal/core/port"
)

type getMatchService struct {
	repository port.MatchRepository
}

func NewGetMatchService(matchRepository port.MatchRepository) port.GetMatchService {
	return &getMatchService{
		repository: matchRepository,
	}
}

func (service getMatchService) GetMatchByID(id string) (*domain.Match, errors.BaseError) {
	match, err := service.repository.FindMatchByID(id)
	if err != nil {
		switch err.(errors.BaseError).GetCode() {
		case defines.MatchNotFoundErrorCode:
			return nil, errors.NewError(defines.MatchNotFoundErrorCode, "match not found")
		default:
			return nil, errors.NewError(defines.UnexpectedErrorCode, "unexpected error when getting match by id")
		}
	}

	return match, nil
}

func (service getMatchService) GetAllMatches() ([]*domain.Match, errors.BaseError) {
	matches, err := service.repository.FindAllMatches()
	if err != nil {
		return nil, errors.NewError(defines.UnexpectedErrorCode, "unexpected error when getting all matches")
	}

	return matches, nil
}
