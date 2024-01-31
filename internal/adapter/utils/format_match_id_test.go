package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatMatchID(t *testing.T) {
	t.Run("should format match id", func(t *testing.T) {
		matchID := "1234"
		expected := "game_1234"
		result := FormatMatchID(matchID)
		assert.Equal(t, expected, result)
	})
}
