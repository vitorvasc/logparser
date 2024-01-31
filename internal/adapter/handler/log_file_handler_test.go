package handler

import (
	"os"
	"reflect"
	"testing"

	"logparser/internal/adapter/dto"
	"logparser/internal/core/domain"
	"logparser/internal/core/port"
	"logparser/internal/core/service"

	"github.com/stretchr/testify/mock"
)

func TestCreateMatchesFromLogFile(t *testing.T) {
	testCases := []struct {
		name        string
		serviceMock port.CreateMatchHistoryService
		inputFile   *os.File
		expected    *dto.ProcessResult
	}{
		{
			name:        "given valid log file, should filter entries and create match history",
			serviceMock: getDefaultServiceMock(),
			inputFile: func() *os.File {
				file, err := os.Open("../../../resources/qgames.log")
				if err != nil {
					panic(err)
				}
				return file
			}(),
			expected: &dto.ProcessResult{
				TotalProcessedMatches: 1,
				Failures:              make([]*dto.ProcessFailure, 0),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			handler := NewLogFileHandler(testCase.serviceMock)
			result := handler.CreateMatchesFromLogFile(testCase.inputFile)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected result %v, got %v", testCase.expected, result)
			}
		})
	}
}

func getDefaultServiceMock() port.CreateMatchHistoryService {
	serviceMock := service.NewCreateMatchHistoryServiceMock()
	serviceMock.On("BulkCreate", mock.Anything).Return([]*domain.BulkCreationResult{
		{
			MatchID:      "3",
			Success:      true,
			ErrorMessage: "",
		},
	}).Once()
	serviceMock.On("Create", mock.Anything).Return(&domain.Match{
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
	}, nil).Once()
	return serviceMock
}
