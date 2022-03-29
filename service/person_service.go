package service

import (
	"bookkeepergo/entity"
	"bookkeepergo/repository"
	"errors"
)

type PersonService struct {
	Repository repository.PersonRepository
}

func (service PersonService) Get(id string) (*entity.Person, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	} else {
		return category, nil
	}
}
