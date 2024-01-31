package service

import (
	"logparser/internal/core/domain"
	"logparser/internal/core/errors"

	"github.com/stretchr/testify/mock"
)

type getMatchServiceMock struct {
	mock.Mock
}

func NewGetMatchServiceMock() *getMatchServiceMock {
	return &getMatchServiceMock{}
}

func (service *getMatchServiceMock) GetMatchByID(matchID string) (*domain.Match, errors.BaseError) {
	args := service.Called(matchID)
	if args.Get(1) != nil {
		return nil, args.Get(1).(errors.BaseError)
	}
	return args.Get(0).(*domain.Match), nil
}
