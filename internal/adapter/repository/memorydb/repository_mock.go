package memorydb

import (
	"logparser/internal/core/domain"

	"github.com/stretchr/testify/mock"
)

type matchRepositoryMock struct {
	mock.Mock
}

func NewMatchRepositoryMock() *matchRepositoryMock {
	return &matchRepositoryMock{}
}

func (repository *matchRepositoryMock) SaveMatch(match *domain.Match) error {
	args := repository.Called(match)
	return args.Error(0)
}

func (repository *matchRepositoryMock) FindMatchByID(id string) (*domain.Match, error) {
	args := repository.Called(id)
	return args.Get(0).(*domain.Match), args.Error(1)
}
