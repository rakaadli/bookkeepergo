package service

import (
	"bookkeepergo/entity"
	"bookkeepergo/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var personRepository = &repository.PersonRepositoryMock{Mock: mock.Mock{}}
var personService = PersonService{Repository: personRepository}

func TestCategoryService_GetNotFound(t *testing.T) {

	// program mock
	personRepository.Mock.On("FindById", "1").Return(nil)

	category, err := personService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryService_GetSuccess(t *testing.T) {
	person := entity.Person{
		Id:   "1",
		Name: "Laptop",
	}
	personRepository.Mock.On("FindById", "2").Return(person)

	result, err := personService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, person.Id, result.Id)
	assert.Equal(t, person.Name, result.Name)
}
