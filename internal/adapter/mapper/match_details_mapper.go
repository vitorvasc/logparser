package mapper

import (
	"sort"

	"logparser/internal/adapter/dto"
	"logparser/internal/core/domain"
)

func FromMatchToMatchDetailsDto(match *domain.Match, shouldGenerateCompleteDetails bool) *dto.MatchDetails {
	matchDetails := new(dto.MatchDetails)
	matchDetails.TotalKills = match.TotalKills
	matchDetails.Players = make([]dto.Player, 0, len(match.Players))
	matchDetails.Kills = make(map[string]int)

	for _, player := range match.Players {
		matchDetails.Players = append(matchDetails.Players, dto.Player(player.Name))
		matchDetails.Kills[player.Name] = player.Kills
	}

	if shouldGenerateCompleteDetails {
		matchDetails.KillsByMeans = make(map[string]int)
		for _, kill := range match.KillHistory {
			matchDetails.KillsByMeans[string(kill.Weapon)]++
		}
	}

	return matchDetails
}

func FromMatchToMatchDetailsMapDto(match *domain.Match, shouldGenerateCompleteDetails bool) map[string]*dto.MatchDetails {
	matchDetailsMap := make(map[string]*dto.MatchDetails)
	matchDetailsMap[match.ID] = FromMatchToMatchDetailsDto(match, shouldGenerateCompleteDetails)
	return matchDetailsMap
}

func FromMatchListToMatchDetailsMapDto(matches []*domain.Match, shouldGenerateCompleteDetails bool) map[string]*dto.MatchDetails {
	matchesMap := make(map[string]*dto.MatchDetails)
	for _, match := range matches {
		matchesMap[match.ID] = FromMatchToMatchDetailsDto(match, shouldGenerateCompleteDetails)
	}

	keys := make([]string, 0, len(matches))
	for key := range matchesMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	sortedMap := make(map[string]*dto.MatchDetails)
	for _, keyValue := range keys {
		sortedMap[keyValue] = matchesMap[keyValue]
	}

	return sortedMap
}
