package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMatch_InsertOrUpdatePlayer(t *testing.T) {
	testCases := []struct {
		name     string
		match    *Match
		player   *Player
		expected []*Player
	}{
		{
			name:     "given a match with no players, should add the player",
			match:    NewMatch("1"),
			player:   NewPlayer(1, "vitor"),
			expected: []*Player{NewPlayer(1, "vitor")},
		},
		{
			name: "given a match with one player, when creating player with same id should update the player",
			match: func() *Match {
				return &Match{
					ID:      "2",
					Players: []*Player{NewPlayer(1, "vitor")},
				}
			}(),
			player: NewPlayer(1, "vitor2"),
			expected: []*Player{
				{
					ID:          1,
					Name:        "vitor2",
					NameHistory: []string{"vitor"},
					Kills:       0,
					Deaths:      0,
				},
			},
		},
		{
			name: "given a match with more than one player, should add the player",
			match: func() *Match {
				return &Match{
					ID: "3",
					Players: []*Player{
						NewPlayer(1, "vitor"),
						NewPlayer(2, "Dono da Bola"),
						NewPlayer(3, "Mocinho"),
					},
				}
			}(),
			player: NewPlayer(4, "Isgalamido"),
			expected: []*Player{
				NewPlayer(1, "vitor"),
				NewPlayer(2, "Dono da Bola"),
				NewPlayer(3, "Mocinho"),
				NewPlayer(4, "Isgalamido"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.match.InsertOrUpdatePlayer(testCase.player)
			require.Equal(t, testCase.expected, testCase.match.Players)
		})
	}
}

func TestMatch_NoticeKill(t *testing.T) {
	testCases := []struct {
		name                      string
		match                     *Match
		kill                      *Kill
		expectedMatchKills        int
		expectedMatchKillsHistory []*Kill
		expectedKillerKills       int
		expectedTargetDeaths      int
	}{
		{
			name: "given a match with no kills, when kill between two players is noticed, should add on total score, show on kill history, add on killer kills and target deaths",
			match: func() *Match {
				return &Match{
					ID:         "3",
					TotalKills: 0,
					Players: []*Player{
						NewPlayer(1, "vitor"),
						NewPlayer(2, "Dono da Bola"),
						NewPlayer(3, "Mocinho"),
					},
					KillHistory: make([]*Kill, 0),
				}
			}(),
			kill:               NewKill(1, "vitor", 2, "Dono da Bola", MOD_ROCKET, 6),
			expectedMatchKills: 1,
			expectedMatchKillsHistory: []*Kill{
				NewKill(1, "vitor", 2, "Dono da Bola", MOD_ROCKET, 6),
			},
			expectedKillerKills:  1,
			expectedTargetDeaths: 1,
		},
		{
			name: "given a match with no kills, when kill by world is noticed, should add on total score, show on kill history, add on target deaths",
			match: func() *Match {
				return &Match{
					ID:         "3",
					TotalKills: 0,
					Players: []*Player{
						NewPlayer(1, "vitor"),
					},
					KillHistory: make([]*Kill, 0),
				}
			}(),
			kill:               NewKill(1022, "<world>", 1, "vitor", MOD_TRIGGER_HURT, 22),
			expectedMatchKills: 1,
			expectedMatchKillsHistory: []*Kill{
				NewKill(1022, "<world>", 1, "vitor", MOD_TRIGGER_HURT, 22),
			},
			expectedKillerKills:  0,
			expectedTargetDeaths: 1,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.match.NoticeKill(testCase.kill)

			require.Equal(t, testCase.expectedMatchKills, testCase.match.TotalKills)
			require.Equal(t, testCase.expectedMatchKillsHistory, testCase.match.KillHistory)
			require.Equal(t, testCase.expectedTargetDeaths, testCase.match.GetPlayerByID(testCase.kill.TargetID).Deaths)

			if !testCase.kill.KillerEqualsWorld() {
				require.Equal(t, testCase.expectedKillerKills, testCase.match.GetPlayerByID(testCase.kill.KillerID).Kills)
			}
		})
	}
}

func TestMatch_GetPlayerByID(t *testing.T) {
	testCases := []struct {
		name     string
		match    *Match
		playerID int
		expected *Player
	}{
		{
			name: "given a match with no players, when getting player by id, should return nil",
			match: func() *Match {
				return &Match{
					ID:      "1",
					Players: []*Player{},
				}
			}(),
			playerID: 1,
			expected: nil,
		},
		{
			name: "given a match with one player, when getting player by id, should return the correct player",
			match: func() *Match {
				return &Match{
					ID:      "2",
					Players: []*Player{NewPlayer(1, "vitor")},
				}
			}(),
			playerID: 1,
			expected: NewPlayer(1, "vitor"),
		},
		{
			name: "given a match with more than one player, when getting player by id, should return the correct player",
			match: func() *Match {
				return &Match{
					ID: "3",
					Players: []*Player{
						NewPlayer(1, "vitor"),
						NewPlayer(2, "Dono da Bola"),
						NewPlayer(3, "Mocinho"),
					},
				}
			}(),
			playerID: 2,
			expected: NewPlayer(2, "Dono da Bola"),
		},
		{
			name: "given a match with more than one player, when getting player that does not exist, should return nil",
			match: func() *Match {
				return &Match{
					ID: "3",
					Players: []*Player{
						NewPlayer(1, "vitor"),
						NewPlayer(2, "Dono da Bola"),
						NewPlayer(3, "Mocinho"),
					},
				}
			}(),
			playerID: 4,
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			require.Equal(t, testCase.expected, testCase.match.GetPlayerByID(testCase.playerID))
		})
	}
}
