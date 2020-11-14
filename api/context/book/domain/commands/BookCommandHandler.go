package commands

import (
	"Modgo/context/book/domain/models"
	"Modgo/context/book/domain/queries"
	"Modgo/context/book/domain/repository"

	"github.com/google/uuid"
)

// BookCommandHandler ...
type BookCommandHandler struct {
	repo repository.IBookRepository
}

// NewBookCommandHandler ...
func NewBookCommandHandler(repo repository.IBookRepository) BookCommandHandler {
	return BookCommandHandler{
		repo: repo,
	}
}

// CreateBookHandle ...
func (c BookCommandHandler) CreateBookHandle(bookCommand CreateBookCommand) error {
	book := models.Book{
		ID:        uuid.New().String(),
		Title:     bookCommand.Title,
		Author:    bookCommand.Author,
		Publisher: bookCommand.Publisher,
	}

	return c.repo.Create(book)
}

// DeleteBookHandle ...
func (c BookCommandHandler) DeleteBookHandle(id string) error {
	return c.repo.Delete(id)
}

// UpdateBookHandle ...
func (c BookCommandHandler) UpdateBookHandle(id string, bookCommand UpdateBookCommand) error {
	book := models.Book{
		ID:        bookCommand.ID,
		Title:     bookCommand.Title,
		Author:    bookCommand.Author,
		Publisher: bookCommand.Publisher,
	}

	return c.repo.Update(id, book)
}

// TODO The methods below need to be segregated in a new QueryHandler.

// GetAllBooksHandle ...
func (c BookCommandHandler) GetAllBooksHandle() ([]queries.BookViewModel, error) {
	b, err := c.repo.GetAll()
	return b, err
}

// GetByIDHandle ...
func (c BookCommandHandler) GetByIDHandle(id string) (queries.BookViewModel, error) {
	b, err := c.repo.GetByID(id)
	return b, err
}
