package mapper

import (
	"logparser/internal/adapter/dto"
	"logparser/internal/core/domain"
)

func FromMatchToMatchDetailsDto(match *domain.Match) *dto.MatchDetails {
	matchDetails := new(dto.MatchDetails)

	matchDetails.TotalKills = match.TotalKills
	matchDetails.Players = make([]dto.Player, 0, len(match.Players))
	matchDetails.Kills = make(map[string]dto.Kill)

	for _, player := range match.Players {
		matchDetails.Players = append(matchDetails.Players, dto.Player(player.Name))
		matchDetails.Kills[player.Name] = dto.Kill(player.Kills)
	}

	return matchDetails
}
