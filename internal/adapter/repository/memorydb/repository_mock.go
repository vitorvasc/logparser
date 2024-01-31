package memorydb

import (
	"logparser/internal/core/domain"
	"logparser/internal/core/errors"

	"github.com/stretchr/testify/mock"
)

type matchRepositoryMock struct {
	mock.Mock
}

func NewMatchRepositoryMock() *matchRepositoryMock {
	return &matchRepositoryMock{}
}

func (repository *matchRepositoryMock) SaveMatch(match *domain.Match) errors.BaseError {
	args := repository.Called(match)
	if args.Get(0) != nil {
		return args.Get(0).(errors.BaseError)
	}
	return nil
}

func (repository *matchRepositoryMock) FindMatchByID(id string) (*domain.Match, errors.BaseError) {
	args := repository.Called(id)
	if args.Get(1) != nil {
		return nil, args.Get(1).(errors.BaseError)
	}
	return args.Get(0).(*domain.Match), nil
}
