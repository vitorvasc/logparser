package handler

import (
	"testing"

	"logparser/internal/adapter/dto"
	"logparser/internal/adapter/utils"
	"logparser/internal/config/defines"
	"logparser/internal/core/domain"
	"logparser/internal/core/errors"
	"logparser/internal/core/port"
	"logparser/internal/core/service"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetMatchHandler_GetSimpleReportByMatchID(t *testing.T) {
	testCases := []struct {
		name          string
		serviceMock   port.GetMatchService
		matchID       string
		expected      map[string]*dto.MatchDetails
		expectedError error
	}{
		{
			name:        "given valid existing match id should return match details",
			serviceMock: getDefaultGetMatchServiceMock(),
			matchID:     "1",
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
			expectedError: nil,
		},
		{
			name: "given invalid match id should return error",
			serviceMock: func() port.GetMatchService {
				serviceMock := service.NewGetMatchServiceMock()
				serviceMock.On("GetMatchByID", mock.Anything).
					Return(nil, errors.NewError(defines.MatchNotFoundErrorCode, "match not found")).
					Once()
				return serviceMock
			}(),
			matchID:       utils.FormatMatchID("2"),
			expected:      nil,
			expectedError: errors.NewError(defines.MatchNotFoundErrorCode, "match not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			handler := NewGetMatchHandler(testCase.serviceMock)
			matchDetails, err := handler.GetSimpleReportByMatchID(testCase.matchID)
			require.Equal(t, testCase.expectedError, err)
			require.Equal(t, testCase.expected, matchDetails)
		})

	}
}

func TestGetMatchHandler_GetCompleteReportByMatchID(t *testing.T) {
	testCases := []struct {
		name          string
		serviceMock   port.GetMatchService
		matchID       string
		expected      map[string]*dto.MatchDetails
		expectedError error
	}{
		{
			name:        "given valid existing match id should return match details",
			serviceMock: getDefaultGetMatchServiceMock(),
			matchID:     "1",
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
			expectedError: nil,
		},
		{
			name: "given invalid match id should return error",
			serviceMock: func() port.GetMatchService {
				serviceMock := service.NewGetMatchServiceMock()
				serviceMock.On("GetMatchByID", mock.Anything).
					Return(nil, errors.NewError(defines.MatchNotFoundErrorCode, "match not found")).
					Once()
				return serviceMock
			}(),
			matchID:       utils.FormatMatchID("2"),
			expected:      nil,
			expectedError: errors.NewError(defines.MatchNotFoundErrorCode, "match not found"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			handler := NewGetMatchHandler(testCase.serviceMock)
			matchDetails, err := handler.GetCompleteReportByMatchID(testCase.matchID)
			require.Equal(t, testCase.expectedError, err)
			require.Equal(t, testCase.expected, matchDetails)
		})

	}
}

func getDefaultGetMatchServiceMock() port.GetMatchService {
	serviceMock := service.NewGetMatchServiceMock()
	serviceMock.On("GetMatchByID", mock.Anything).Return(&domain.Match{
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
	}, nil).Once()
	return serviceMock
}
