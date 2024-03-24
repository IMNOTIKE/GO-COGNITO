package entities

import (
	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID
	ISBN      string
	Title     string
	Author    string
	Year      int
	Genre     string
	Status    string
	Publisher string
	StartDate string
}
