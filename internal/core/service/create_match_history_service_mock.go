package service

import (
	"logparser/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type createMatchHistoryServiceMock struct {
	mock.Mock
}

func NewCreateMatchHistoryServiceMock() *createMatchHistoryServiceMock {
	return &createMatchHistoryServiceMock{}
}

func (service *createMatchHistoryServiceMock) BulkCreate(matchHistoryList []*domain.MatchHistory) []*domain.BulkCreationResult {
	args := service.Called(matchHistoryList)
	return args.Get(0).([]*domain.BulkCreationResult)
}

func (service *createMatchHistoryServiceMock) Create(matchHistory *domain.MatchHistory) (*domain.Match, error) {
	args := service.Called(matchHistory)
	return args.Get(0).(*domain.Match), args.Error(1)
}
