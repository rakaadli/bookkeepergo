package repository

import (
	"bookkeepergo/entity"
	"github.com/stretchr/testify/mock"
)

type PersonRepositoryMock struct {
	Mock mock.Mock
}

func (repository *PersonRepositoryMock) FindById(id string) *entity.Person {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Person)
		return &category
	}
}
