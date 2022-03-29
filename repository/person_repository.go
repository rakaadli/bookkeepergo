package repository

import "book_keeper_go/entity"

type PersonRepository interface {
	FindById(id string) *entity.Person
}
