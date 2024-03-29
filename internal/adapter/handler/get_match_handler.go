package handler

import (
	"logparser/internal/adapter/dto"
	"logparser/internal/adapter/mapper"
	"logparser/internal/adapter/utils"
	"logparser/internal/core/port"
)

type GetMatchHandler struct {
	service port.GetMatchService
}

func NewGetMatchHandler(getMatchService port.GetMatchService) GetMatchHandler {
	return GetMatchHandler{
		service: getMatchService,
	}
}

func (h GetMatchHandler) GetSimpleReportByMatchID(matchID string) (map[string]*dto.MatchDetails, error) {
	match, err := h.service.GetMatchByID(utils.FormatMatchID(matchID))
	if err != nil {
		return nil, err
	}

	return mapper.FromMatchToMatchDetailsMapDto(match, false), nil
}

func (h GetMatchHandler) GetCompleteReportByMatchID(matchID string) (map[string]*dto.MatchDetails, error) {
	match, err := h.service.GetMatchByID(utils.FormatMatchID(matchID))
	if err != nil {
		return nil, err
	}

	return mapper.FromMatchToMatchDetailsMapDto(match, true), nil
}

func (h GetMatchHandler) GetCompleteReportForAllMatches() (map[string]*dto.MatchDetails, error) {
	matches, err := h.service.GetAllMatches()
	if err != nil {
		return nil, err
	}

	return mapper.FromMatchListToMatchDetailsMapDto(matches, true), nil
}
