package domain

import (
	"testing"

	"logparser/internal/config/defines"

	"github.com/stretchr/testify/require"
)

func TestKill_KillerEqualsWorld(t *testing.T) {
	testCases := []struct {
		name          string
		kill          *Kill
		expectedValue bool
	}{
		{
			name:          "given a kill with world as killer, should return true",
			kill:          NewKill(1022, defines.WorldPlayerName, 1, "Vitor", MOD_TRIGGER_HURT, 22),
			expectedValue: true,
		},
		{
			name:          "given a kill with player as killer, should return false",
			kill:          NewKill(3, "Dono da Bola", 1, "Vitor", MOD_RAILGUN, 10),
			expectedValue: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.kill.KillerEqualsWorld()
			require.Equal(t, testCase.expectedValue, result)
		})
	}
}

func TestKill_KillerEqualsTarget(t *testing.T) {
	testCases := []struct {
		name          string
		kill          *Kill
		expectedValue bool
	}{
		{
			name:          "given a kill where player killed itself, should return true",
			kill:          NewKill(1, "Vitor", 1, "Vitor", MOD_ROCKET, 6),
			expectedValue: true,
		},
		{
			name:          "given a kill with different player as killer, should return false",
			kill:          NewKill(3, "Dono da Bola", 1, "Vitor", MOD_RAILGUN, 10),
			expectedValue: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := testCase.kill.KillerEqualsTarget()
			require.Equal(t, testCase.expectedValue, result)
		})
	}
}
