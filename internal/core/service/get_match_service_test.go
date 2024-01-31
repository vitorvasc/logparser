package service

import (
	"testing"

	"logparser/internal/adapter/repository/memorydb"
	"logparser/internal/config/defines"
	"logparser/internal/core/domain"
	"logparser/internal/core/errors"
	"logparser/internal/core/port"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetMatchService_GetMatchByID(t *testing.T) {
	testCases := []struct {
		name            string
		matchRepository port.MatchRepository
		expectedResult  *domain.Match
		expectedError   errors.BaseError
	}{
		{
			name: "should return match when match exists",
			matchRepository: func() port.MatchRepository {
				matchRepository := memorydb.NewMatchRepositoryMock()
				matchRepository.
					On("FindMatchByID", mock.Anything).
					Return(&domain.Match{
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
					}, nil).
					Once()
				return matchRepository
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
			name: "should return not found error if match does not exist",
			matchRepository: func() port.MatchRepository {
				matchRepository := memorydb.NewMatchRepositoryMock()
				matchRepository.
					On("FindMatchByID", mock.Anything).
					Return(nil, errors.NewError(defines.MatchNotFoundErrorCode, "match not found")).
					Once()
				return matchRepository
			}(),
			expectedResult: nil,
			expectedError:  errors.NewError(defines.MatchNotFoundErrorCode, "match not found"),
		},
		{
			name: "should return unexpected error if repository returns unexpected error",
			matchRepository: func() port.MatchRepository {
				matchRepository := memorydb.NewMatchRepositoryMock()
				matchRepository.
					On("FindMatchByID", mock.Anything).
					Return(nil, errors.NewError(defines.UnexpectedErrorCode, "unexpected error when getting match by id")).
					Once()
				return matchRepository
			}(),
			expectedResult: nil,
			expectedError:  errors.NewError(defines.UnexpectedErrorCode, "unexpected error when getting match by id"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := NewGetMatchService(testCase.matchRepository)
			result, err := service.GetMatchByID("3")

			require.Equal(t, testCase.expectedResult, result)
			require.Equal(t, testCase.expectedError, err)
		})
	}
}

func TestGetMatchService_GetAllMatches(t *testing.T) {
	testCases := []struct {
		name            string
		matchRepository port.MatchRepository
		expectedResult  []*domain.Match
		expectedError   errors.BaseError
	}{
		{
			name: "should return all matches when matches exist",
			matchRepository: func() port.MatchRepository {
				matchRepository := memorydb.NewMatchRepositoryMock()
				matchRepository.
					On("FindAllMatches", mock.Anything).
					Return([]*domain.Match{
						{
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
						{
							ID:         "4",
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
					}, nil).
					Once()
				return matchRepository
			}(),
			expectedResult: []*domain.Match{
				{
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
				{
					ID:         "4",
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
			},
			expectedError: nil,
		},
		{
			name: "should return empty list when no matches exist",
			matchRepository: func() port.MatchRepository {
				matchRepository := memorydb.NewMatchRepositoryMock()
				matchRepository.
					On("FindAllMatches", mock.Anything).
					Return(make([]*domain.Match, 0), nil).
					Once()
				return matchRepository
			}(),
			expectedResult: make([]*domain.Match, 0),
			expectedError:  nil,
		},
		{
			name: "should return unexpected error if repository returns unexpected error",
			matchRepository: func() port.MatchRepository {
				matchRepository := memorydb.NewMatchRepositoryMock()
				matchRepository.
					On("FindAllMatches", mock.Anything).
					Return(nil, errors.NewError(defines.UnexpectedErrorCode, "unexpected error when getting all matches")).
					Once()
				return matchRepository
			}(),
			expectedResult: nil,
			expectedError:  errors.NewError(defines.UnexpectedErrorCode, "unexpected error when getting all matches"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := NewGetMatchService(testCase.matchRepository)
			result, err := service.GetAllMatches()

			require.Equal(t, testCase.expectedResult, result)
			require.Equal(t, testCase.expectedError, err)
		})
	}
}
