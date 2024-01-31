package service

import (
	"logparser/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type createMatchServiceMock struct {
	mock.Mock
}

func NewCreateMatchServiceMock() *createMatchServiceMock {
	return &createMatchServiceMock{}
}

func (service *createMatchServiceMock) BulkCreate(matchHistoryList []*domain.MatchHistory) []*domain.BulkCreationResult {
	args := service.Called(matchHistoryList)
	return args.Get(0).([]*domain.BulkCreationResult)
}

func (service *createMatchServiceMock) Create(matchHistory *domain.MatchHistory) (*domain.Match, error) {
	args := service.Called(matchHistory)
	return args.Get(0).(*domain.Match), args.Error(1)
}
