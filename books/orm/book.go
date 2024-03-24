package orm

import (
	"fmt"

	. "cognito/books/config"
	. "cognito/books/entities"

	"github.com/go-pg/pg/v10"
	"github.com/gofrs/uuid/v5"
)

var (
	bookConfig AppConfig
	dbUrl      string
	book       Book
	books      []Book
)

func Connect(config AppConfig) *pg.DB {
	bookConfig = config
	return pg.Connect(&pg.Options{
		User:     bookConfig.Db.User,
		Password: bookConfig.Db.Pwd,
		Addr:     fmt.Sprintf("%s:%d", bookConfig.Db.Url, bookConfig.Db.Port),
		Database: bookConfig.Db.DbName,
		PoolSize: bookConfig.Db.PoolSize,
	})
}

func GetBooks(pg *pg.DB) ([]Book, error) {

	err := pg.Model(&books).Select()
	return books, err

}

func GetBookById(pg *pg.DB, id uuid.UUID) (Book, error) {

	err := pg.Model(&book).Where("id = ?0", id).Select()
	return book, err

}

func GetBooksByISBN(pg *pg.DB, isbn string) ([]Book, error) {
	err := pg.Model(&books).Where("isbn = ?0", isbn).Select()
	return books, err
}

func CreateBook(pg *pg.DB, newBook Book) error {
	_, err := pg.Model(&newBook).Insert()
	return err
}

func UpdateBook(pg *pg.DB, updatedBook Book) error {
	_, err := pg.Model(&updatedBook).Where("id = ?0", updatedBook.Id).Update()
	return err
}

func DeleteBook(pg *pg.DB, id uuid.UUID) error {
	book.Id = id
	_, err := pg.Model(&book).Where("id = ?0", id).Delete()
	return err
}
