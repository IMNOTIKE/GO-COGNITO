package service

import (
	"cognito/books/entities"
)

func GetAll() []entities.Book {

	return make([]entities.Book, 0)
}
