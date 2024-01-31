package memorydb

import (
	"testing"

	"logparser/internal/config/defines"
	"logparser/internal/core/domain"
	apperrors "logparser/internal/core/errors"

	"github.com/stretchr/testify/require"
)

func TestSaveMatch(t *testing.T) {
	r := NewMatchRepository()

	match := &domain.Match{
		ID: "123",
	}

	err := r.SaveMatch(match)

	require.Nil(t, err)
}

func TestFindMatchByID(t *testing.T) {
	testCases := []struct {
		name          string
		id            string
		expectedError error
		expectedMatch *domain.Match
	}{
		{
			name:          "given valid match, should return match found",
			id:            "123",
			expectedError: nil,
			expectedMatch: &domain.Match{
				ID: "123",
			},
		},
		{
			name:          "given invalid match id, should return match not found",
			id:            "12345",
			expectedError: apperrors.NewError(defines.MatchNotFoundErrorCode, "match not found"),
			expectedMatch: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			r := NewMatchRepository()

			err := r.SaveMatch(&domain.Match{
				ID: "123",
			})
			require.Nil(t, err)

			match, err := r.FindMatchByID(testCase.id)
			require.Equal(t, testCase.expectedError, err)
			require.Equal(t, testCase.expectedMatch, match)
		})
	}
}

func TestFindAllMatches(t *testing.T) {
	r := NewMatchRepository()

	err := r.SaveMatch(&domain.Match{
		ID: "123",
	})
	require.Nil(t, err)

	err = r.SaveMatch(&domain.Match{
		ID: "456",
	})
	require.Nil(t, err)

	matchList, err := r.FindAllMatches()
	require.Nil(t, err)
	require.Len(t, matchList, 2)
	require.Equal(t, "123", matchList[0].ID)
	require.Equal(t, "456", matchList[1].ID)
}
