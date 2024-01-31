package domain

import (
	"testing"
)

func TestPlayer_ChangeName(t *testing.T) {
	testCases := []struct {
		name                string
		player              *Player
		newName             string
		expectedNameHistory []string
	}{
		{
			name:                "given a player with the same name, should do nothing",
			player:              NewPlayer(1, "vitor"),
			newName:             "vitor",
			expectedNameHistory: []string{},
		},
		{
			name:                "given a player with a different name, should change it and append the old name to the history",
			player:              NewPlayer(1, "vitor"),
			newName:             "vitor2",
			expectedNameHistory: []string{"vitor"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.player.ChangeName(testCase.newName)
			if testCase.player.Name != testCase.newName {
				t.Errorf("expected name %s, got %s", testCase.newName, testCase.player.Name)
			}
			if len(testCase.player.NameHistory) != len(testCase.expectedNameHistory) {
				t.Errorf("expected history length %v, got %v", len(testCase.expectedNameHistory), len(testCase.player.NameHistory))
			}
		})
	}
}

func TestPlayer_AddKill(t *testing.T) {
	testCases := []struct {
		name     string
		player   *Player
		expected int
	}{
		{
			name:     "given a player, should add one kill",
			player:   NewPlayer(1, "vitor"),
			expected: 1,
		},
		{
			name: "given a player with one kill, should add one more kill",
			player: func() *Player {
				return &Player{
					ID:    1,
					Name:  "vitor",
					Kills: 1,
				}
			}(),
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.player.AddKill()
			if testCase.player.Kills != testCase.expected {
				t.Errorf("expected %v kills, got %v", testCase.expected, testCase.player.Kills)
			}
		})
	}
}

func TestPlayer_AddDeath(t *testing.T) {
	testCases := []struct {
		name     string
		player   *Player
		expected int
	}{
		{
			name:     "given a player, should add one death",
			player:   NewPlayer(1, "vitor"),
			expected: 1,
		},
		{
			name: "given a player with one death, should add one more death",
			player: func() *Player {
				return &Player{
					ID:     1,
					Name:   "vitor",
					Deaths: 1,
				}
			}(),
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.player.AddDeath()
			if testCase.player.Deaths != testCase.expected {
				t.Errorf("expected %v deaths, got %v", testCase.expected, testCase.player.Deaths)
			}
		})
	}
}

func TestPlayer_Equals(t *testing.T) {
	testCases := []struct {
		name     string
		player   *Player
		other    *Player
		expected bool
	}{
		{
			name: "given two players with the same id, should return true",
			player: func() *Player {
				return &Player{
					ID: 1,
				}
			}(),
			other: func() *Player {
				return &Player{
					ID: 1,
				}
			}(),
			expected: true,
		},
		{
			name: "given two players with different ids, should return false",
			player: func() *Player {
				return &Player{
					ID: 1,
				}
			}(),
			other: func() *Player {
				return &Player{
					ID: 2,
				}
			}(),
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.player.Equals(testCase.other) != testCase.expected {
				t.Errorf("expected %v, got %v", testCase.expected, testCase.player.Equals(testCase.other))
			}
		})
	}
}
