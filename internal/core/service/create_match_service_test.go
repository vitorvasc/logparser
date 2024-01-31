package service

import (
	"testing"

	"logparser/internal/adapter/repository/memorydb"
	"logparser/internal/config/defines"
	"logparser/internal/core/domain"
	apperrors "logparser/internal/core/errors"
	"logparser/internal/core/port"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateMatchHistoryService_Create(t *testing.T) {
	testCases := []struct {
		name            string
		matchRepository port.MatchRepository
		matchHistory    *domain.MatchHistory
		expectedResult  *domain.Match
		expectedError   error
	}{
		{
			name:            "given a match history with valid logs, should return a created match",
			matchRepository: loadDefaultMatchRepository(),
			matchHistory: func() *domain.MatchHistory {
				return &domain.MatchHistory{
					ID: "1",
					Logs: []*domain.Log{
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "21:15 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\uriel/zael\\hmodel\\uriel/zael\\g_redteam\\\\g_blueteam\\\\c1\\5\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0",
						},
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "21:17 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\uriel/zael\\hmodel\\uriel/zael\\g_redteam\\\\g_blueteam\\\\c1\\5\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0",
						},
						{
							Type:  defines.LogTypeKill,
							Value: "21:42 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
						},
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "21:51 ClientUserinfoChanged: 3 n\\Dono da Bola\\t\\0\\model\\sarge/krusade\\hmodel\\sarge/krusade\\g_redteam\\\\g_blueteam\\\\c1\\5\\c2\\5\\hc\\95\\w\\0\\l\\0\\tt\\0\\tl\\0",
						},
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "21:53 ClientUserinfoChanged: 3 n\\Mocinha\\t\\0\\model\\sarge\\hmodel\\sarge\\g_redteam\\\\g_blueteam\\\\c1\\4\\c2\\5\\hc\\95\\w\\0\\l\\0\\tt\\0\\tl\\0",
						},
						{
							Type:  defines.LogTypeKill,
							Value: "22:06 Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH",
						},
						{
							Type:  defines.LogTypeKill,
							Value: "22:18 Kill: 2 2 7: Isgalamido killed Isgalamido by MOD_ROCKET_SPLASH",
						},
						{
							Type:  defines.LogTypeKill,
							Value: "23:06 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
						},
						{
							Type:  defines.LogTypeKill,
							Value: "29:08 Kill: 3 2 6: Mocinha killed Isgalamido by MOD_ROCKET",
						},
					},
				}
			}(),
			expectedResult: &domain.Match{
				ID:         "1",
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
			expectedError: nil,
		},
		{
			name:            "given a match history with unknown logs, should return an empty match",
			matchRepository: loadDefaultMatchRepository(),
			matchHistory: func() *domain.MatchHistory {
				return &domain.MatchHistory{
					ID: "2",
					Logs: []*domain.Log{
						{
							Type:  defines.LogTypeUnknown,
							Value: "1:51 Item: 3 item_armor_shard",
						},
						{
							Type:  defines.LogTypeUnknown,
							Value: "1:51 Item: 3 item_armor_shard",
						},
						{
							Type:  defines.LogTypeUnknown,
							Value: "1:51 Item: 3 item_armor_shard",
						},
						{
							Type:  defines.LogTypeUnknown,
							Value: "1:51 Item: 3 item_armor_combat",
						},
					},
				}
			}(),
			expectedResult: &domain.Match{
				ID:          "2",
				TotalKills:  0,
				Players:     make([]*domain.Player, 0),
				KillHistory: make([]*domain.Kill, 0),
			},
			expectedError: nil,
		},
		{
			name:            "given a match history with valid logs, a connected player and no kills, should return a match with single player and no kills",
			matchRepository: loadDefaultMatchRepository(),
			matchHistory: func() *domain.MatchHistory {
				return &domain.MatchHistory{
					ID: "3",
					Logs: []*domain.Log{
						{
							Type:  defines.LogTypeStartMatch,
							Value: "0:00 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0",
						},
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "20:34 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\xian/default\\hmodel\\xian/default\\g_redteam\\\\g_blueteam\\\\c1\\4\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0",
						},
						{
							Type:  defines.LogTypeEndMatch,
							Value: "20:37 ShutdownGame:",
						},
					},
				}
			}(),
			expectedResult: &domain.Match{
				ID:         "3",
				TotalKills: 0,
				Players: []*domain.Player{
					{
						ID:          2,
						Name:        "Isgalamido",
						NameHistory: make([]string, 0),
						Kills:       0,
						Deaths:      0,
					},
				},
				KillHistory: make([]*domain.Kill, 0),
			},
			expectedError: nil,
		},
		{
			name:            "given a match history with invalid kill info on logs, when fails on parsing regex should return an error",
			matchRepository: loadDefaultMatchRepository(),
			matchHistory: func() *domain.MatchHistory {
				return &domain.MatchHistory{
					ID: "4",
					Logs: []*domain.Log{
						{
							Type:  defines.LogTypeKill,
							Value: "INVALID_LOG",
						},
					},
				}
			}(),
			expectedResult: nil,
			expectedError:  apperrors.NewError(defines.GetKillInfoErrorCode, "error parsing kill info: INVALID_LOG"),
		},
		{
			name:            "given a match history with invalid player connection info on logs, when fails on parsing regex should return an error",
			matchRepository: loadDefaultMatchRepository(),
			matchHistory: func() *domain.MatchHistory {
				return &domain.MatchHistory{
					ID: "5",
					Logs: []*domain.Log{
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "INVALID_LOG",
						},
					},
				}
			}(),
			expectedResult: nil,
			expectedError:  apperrors.NewError(defines.GetPlayerInfoErrorCode, "error parsing player info: INVALID_LOG"),
		},
		{
			name: "given a match history with valid logs, when save on repository fails should return an error",
			matchRepository: func() port.MatchRepository {
				matchRepository := memorydb.NewMatchRepositoryMock()
				matchRepository.
					On("SaveMatch", mock.Anything).
					Return(apperrors.NewError(defines.SaveMatchErrorCode, "error on save match")).
					Once()
				return matchRepository
			}(),
			matchHistory: func() *domain.MatchHistory {
				return &domain.MatchHistory{
					ID: "6",
					Logs: []*domain.Log{
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "20:34 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\xian/default\\hmodel\\xian/default\\g_redteam\\\\g_blueteam\\\\c1\\4\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0",
						},
					},
				}
			}(),
			expectedResult: nil,
			expectedError:  apperrors.NewError(defines.SaveMatchErrorCode, "error on save match"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			matchService := NewCreateMatchService(testCase.matchRepository)
			result, err := matchService.Create(testCase.matchHistory)
			require.Equal(t, testCase.expectedResult, result)
			require.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestCreateMatchHistoryService_BulkCreate(t *testing.T) {
	testCases := []struct {
		name             string
		matchRepository  port.MatchRepository
		matchHistoryList []*domain.MatchHistory
		expectedResult   []*domain.BulkCreationResult
	}{
		{
			name:            "given a match history with valid logs, should return a list of created matches and success status for all",
			matchRepository: loadDefaultMatchRepository(),
			matchHistoryList: []*domain.MatchHistory{
				{
					ID: "3",
					Logs: []*domain.Log{
						{
							Type:  defines.LogTypeStartMatch,
							Value: "0:00 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0",
						},
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "20:34 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\xian/default\\hmodel\\xian/default\\g_redteam\\\\g_blueteam\\\\c1\\4\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0",
						},
						{
							Type:  defines.LogTypeEndMatch,
							Value: "20:37 ShutdownGame:",
						},
					},
				},
			},
			expectedResult: []*domain.BulkCreationResult{
				{
					MatchID:      "3",
					Success:      true,
					ErrorMessage: "",
				},
			},
		},
		{
			name:            "given a match history with both valid and invalid logs, should return a list of created matches and success and error statuses",
			matchRepository: loadDefaultMatchRepository(),
			matchHistoryList: []*domain.MatchHistory{
				{
					ID: "4",
					Logs: []*domain.Log{
						{
							Type:  defines.LogTypeClientUserInfoChanged,
							Value: "INVALID_LOG",
						},
					},
				},
			},
			expectedResult: []*domain.BulkCreationResult{
				{
					MatchID:      "4",
					Success:      false,
					ErrorMessage: "error parsing player info: INVALID_LOG",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			matchService := NewCreateMatchService(testCase.matchRepository)
			result := matchService.BulkCreate(testCase.matchHistoryList)
			require.Equal(t, testCase.expectedResult, result)
		})
	}
}

func loadDefaultMatchRepository() port.MatchRepository {
	matchRepository := memorydb.NewMatchRepositoryMock()
	matchRepository.
		On("SaveMatch", mock.Anything).
		Return(nil).
		Once()
	return matchRepository
}
