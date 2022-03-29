package service

import (
	"book_keeper_go/entity"
	"book_keeper_go/repository"
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
	category := entity.Person{
		Id:   "1",
		Name: "Laptop",
	}
	personRepository.Mock.On("FindById", "2").Return(category)

	result, err := personService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
	assert.Equal(t, category.Name, result.Name)
}
