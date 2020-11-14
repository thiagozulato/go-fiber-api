package infrasctructure

import (
	"Modgo/context/book/domain/models"
	"Modgo/context/book/domain/queries"
	"database/sql"
)

// BookRepository ...
type BookRepository struct {
	db    *sql.DB
	books []models.Book
}

// NewBookRepository ...
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db:    db,
		books: []models.Book{},
	}
}

// GetAll ...
func (r BookRepository) GetAll() ([]queries.BookViewModel, error) {
	return modelToViewModels(r.books), nil
}

// GetByID ...
func (r BookRepository) GetByID(id string) (queries.BookViewModel, error) {
	book, _ := findBook(id, r.books)
	return modelToViewModel(book), nil
}

// Create ...
func (r *BookRepository) Create(book models.Book) error {
	r.books = append(r.books, book)
	return nil
}

// Update ...
func (r *BookRepository) Update(id string, book models.Book) error {
	_, i := findBook(id, r.books)

	r.books[i] = book

	return nil
}

// Delete ...
func (r *BookRepository) Delete(id string) error {
	_, i := findBook(id, r.books)
	r.books = append(r.books[:i], r.books[i+1:]...)
	return nil
}

func findBook(id string, books []models.Book) (bookmodel models.Book, index int) {
	for i, book := range books {
		if book.ID == id {
			bookmodel = book
			index = i
			break
		}
	}

	return
}

func modelToViewModel(book models.Book) queries.BookViewModel {
	return queries.BookViewModel{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Publisher: book.Publisher,
	}
}

func modelToViewModels(books []models.Book) []queries.BookViewModel {
	var bViewModels []queries.BookViewModel

	for _, book := range books {
		bViewModels = append(bViewModels, queries.BookViewModel{
			ID:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			Publisher: book.Publisher,
		})
	}

	return bViewModels
}
