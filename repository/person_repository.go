package repository

import "bookkeepergo/entity"

type PersonRepository interface {
	FindById(id string) *entity.Person
}
