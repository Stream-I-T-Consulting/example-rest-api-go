package repositories

import (
	"stream-it-consulting/innovation-team/example-rest-api/models"
)

type (
	IBookRepository interface {
		GetBooks() ([]models.Book, error)
		GetBookByID(id int) (*models.Book, error)
		CreateBook(book *models.Book) (*models.Book, error)
		// UpdateBook(book *models.Book) (*models.Book, error)
		// DeleteBook(id int) error
	}
)
