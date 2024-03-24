package entities

import (
	"github.com/gofrs/uuid/v5"
)

type Book struct {
	tableName struct{}  `pg:"cognito-books"`
	Id        uuid.UUID `pg:"id,pk"`
	ISBN      string    `pg:"isbn"`
	Title     string    `pg:"title"`
	Author    string    `pg:"author"`
	Year      int       `pg:"year"`
	Genre     string    `pg:"genre"`
	Status    string    `pg:"status"`
	Publisher string    `pg:"publisher"`
	StartDate string    `pg:"start_date"`
}
