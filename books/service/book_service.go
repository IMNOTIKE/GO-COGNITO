package service

import (
	"cognito/books/entities"
	"cognito/books/orm"
	"errors"

	"github.com/gofrs/uuid/v5"
)

func GetAll() ([]entities.Book, error) {
	return orm.GetBooks(db)
}

func GetById(id uuid.UUID) (entities.Book, error) {
	if id == uuid.Nil {
		return entities.Book{}, errors.New("Nill uuid")
	}
	return orm.GetBookById(db, id)
}

func GetByISBN(isbn string) ([]entities.Book, error) {
	if isbn == "" {
		return []entities.Book{}, errors.New("Empty isbn")
	}
	return orm.GetBooksByISBN(db, isbn)
}

func Create(newBook entities.Book) error {
	return orm.CreateBook(db, newBook)
}

func Update(updatedBook entities.Book) error {
	return orm.UpdateBook(db, updatedBook)
}

func Delete(id uuid.UUID) error {
	return orm.DeleteBook(db, id)
}
