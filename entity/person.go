package entity

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Person struct {
	gorm.Model

	Id    string
	Name  string
	Email string `gorm:"typevarchar(100);unique_index"`
	//Books []Book
}

//type Book struct {
//	gorm.Model
//
//	Title      string
//	Author     string
//	CallNumber int
//	PersonID   int
//}
