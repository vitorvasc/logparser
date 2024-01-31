package mapper

import (
	"testing"

	"logparser/internal/adapter/dto"
	"logparser/internal/core/domain"

	"github.com/stretchr/testify/require"
)

func TestFromMatchToMatchDetailsDto(t *testing.T) {
	testCases := []struct {
		name                  string
		match                 *domain.Match
		expectCompleteDetails bool
		expected              map[string]*dto.MatchDetails
	}{
		{
			name: "given valid match, when mapping match details without generating complete details, should return simple match details dto",
			match: &domain.Match{
				ID:         "game_1",
				TotalKills: 5,
				Players: []*domain.Player{
					{
						ID:          2,
						Name:        "Isgalamido",
						NameHistory: make([]string, 0),
						Kills:       1,
						Deaths:      4,
					},
					{
						ID:   3,
						Name: "Mocinha",
						NameHistory: []string{
							"Dono da Bola",
						},
						Kills:  1,
						Deaths: 1,
					},
					{
						ID:   4,
						Name: "Zeh",
						NameHistory: []string{
							"Vitor",
						},
						Kills:  0,
						Deaths: 0,
					},
				},
				KillHistory: []*domain.Kill{
					{
						KillerID:   1022,
						KillerName: "<world>",
						TargetID:   2,
						TargetName: "Isgalamido",
						Weapon:     domain.MOD_TRIGGER_HURT,
						WeaponID:   22,
					},
					{
						KillerID:   2,
						KillerName: "Isgalamido",
						TargetID:   3,
						TargetName: "Mocinha",
						Weapon:     domain.MOD_ROCKET_SPLASH,
						WeaponID:   7,
					},
					{
						KillerID:   2,
						KillerName: "Isgalamido",
						TargetID:   2,
						TargetName: "Isgalamido",
						Weapon:     domain.MOD_ROCKET_SPLASH,
						WeaponID:   7,
					},
					{
						KillerID:   1022,
						KillerName: "<world>",
						TargetID:   2,
						TargetName: "Isgalamido",
						Weapon:     domain.MOD_TRIGGER_HURT,
						WeaponID:   22,
					},
					{
						KillerID:   3,
						KillerName: "Mocinha",
						TargetID:   2,
						TargetName: "Isgalamido",
						Weapon:     domain.MOD_ROCKET,
						WeaponID:   6,
					},
				},
			},
			expectCompleteDetails: false,
			expected: map[string]*dto.MatchDetails{
				"game_1": {
					TotalKills: 5,
					Players: []dto.Player{
						"Isgalamido",
						"Mocinha",
						"Zeh",
					},
					Kills: map[string]int{
						"Isgalamido": 1,
						"Mocinha":    1,
						"Zeh":        0,
					},
				},
			},
		},
		{
			name: "given valid match, when mapping match details requesting complete details, should return complete match details dto",
			match: &domain.Match{
				ID:         "game_1",
				TotalKills: 5,
				Players: []*domain.Player{
					{
						ID:          2,
						Name:        "Isgalamido",
						NameHistory: make([]string, 0),
						Kills:       1,
						Deaths:      4,
					},
					{
						ID:   3,
						Name: "Mocinha",
						NameHistory: []string{
							"Dono da Bola",
						},
						Kills:  1,
						Deaths: 1,
					},
					{
						ID:   4,
						Name: "Zeh",
						NameHistory: []string{
							"Vitor",
						},
						Kills:  0,
						Deaths: 0,
					},
				},
				KillHistory: []*domain.Kill{
					{
						KillerID:   1022,
						KillerName: "<world>",
						TargetID:   2,
						TargetName: "Isgalamido",
						Weapon:     domain.MOD_TRIGGER_HURT,
						WeaponID:   22,
					},
					{
						KillerID:   2,
						KillerName: "Isgalamido",
						TargetID:   3,
						TargetName: "Mocinha",
						Weapon:     domain.MOD_ROCKET_SPLASH,
						WeaponID:   7,
					},
					{
						KillerID:   2,
						KillerName: "Isgalamido",
						TargetID:   2,
						TargetName: "Isgalamido",
						Weapon:     domain.MOD_ROCKET_SPLASH,
						WeaponID:   7,
					},
					{
						KillerID:   1022,
						KillerName: "<world>",
						TargetID:   2,
						TargetName: "Isgalamido",
						Weapon:     domain.MOD_TRIGGER_HURT,
						WeaponID:   22,
					},
					{
						KillerID:   3,
						KillerName: "Mocinha",
						TargetID:   2,
						TargetName: "Isgalamido",
						Weapon:     domain.MOD_ROCKET,
						WeaponID:   6,
					},
				},
			},
			expectCompleteDetails: true,
			expected: map[string]*dto.MatchDetails{
				"game_1": {
					TotalKills: 5,
					Players: []dto.Player{
						"Isgalamido",
						"Mocinha",
						"Zeh",
					},
					Kills: map[string]int{
						"Isgalamido": 1,
						"Mocinha":    1,
						"Zeh":        0,
					},
					KillsByMeans: map[string]int{
						"MOD_ROCKET":        1,
						"MOD_ROCKET_SPLASH": 2,
						"MOD_TRIGGER_HURT":  2,
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			matchDetails := FromMatchToMatchDetailsDto(testCase.match, testCase.expectCompleteDetails)

			require.Equal(t, testCase.expected, matchDetails)

			// require.Equal(t, testCase.expected.TotalKills, matchDetails[testCase.match.ID].TotalKills)
			// require.Equal(t, len(testCase.expected.Players), len(matchDetails[testCase.match.ID].Players))
		})
	}
}
