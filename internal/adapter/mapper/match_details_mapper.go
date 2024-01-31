package mapper

import (
	"logparser/internal/adapter/dto"
	"logparser/internal/core/domain"
)

func FromMatchToMatchDetailsDto(match *domain.Match, shouldGenerateCompleteDetails bool) map[string]*dto.MatchDetails {
	matchDetailsMap := make(map[string]*dto.MatchDetails)
	matchID := match.ID

	matchDetailsMap[matchID] = new(dto.MatchDetails)
	matchDetailsMap[matchID].TotalKills = match.TotalKills
	matchDetailsMap[matchID].Players = make([]dto.Player, 0, len(match.Players))
	matchDetailsMap[matchID].Kills = make(map[string]int)

	for _, player := range match.Players {
		matchDetailsMap[matchID].Players = append(matchDetailsMap[matchID].Players, dto.Player(player.Name))
		matchDetailsMap[matchID].Kills[player.Name] = player.Kills
	}

	if shouldGenerateCompleteDetails {
		matchDetailsMap[matchID].KillsByMeans = make(map[string]int)
		for _, kill := range match.KillHistory {
			matchDetailsMap[matchID].KillsByMeans[string(kill.Weapon)]++
		}
	}

	return matchDetailsMap
}
