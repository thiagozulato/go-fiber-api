package repository

import (
	"Modgo/context/book/domain/models"
	"Modgo/context/book/domain/queries"
)

// IBookRepository ...
type IBookRepository interface {
	GetAll() ([]queries.BookViewModel, error)
	GetByID(id string) (queries.BookViewModel, error)
	Create(book models.Book) error
	Update(id string, book models.Book) error
	Delete(id string) error
}
