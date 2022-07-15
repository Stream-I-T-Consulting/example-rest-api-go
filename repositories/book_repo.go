package repositories

import (
	"database/sql"
	"stream-it-consulting/innovation-team/example-rest-api/models"
)

type (
	bookRepository struct {
		db *sql.DB
	}
)

// NewBookRepository creates a new book repository
func NewBookRepository(dbConn *sql.DB) IBookRepository {
	return bookRepository{db: dbConn}
}

// CreateBook creates a new book
func (r bookRepository) CreateBook(book *models.Book) (*models.Book, error) {
	// Raw SQL query
	query := `INSERT INTO books 
						(title, author, year) 
						VALUES ($1, $2, $3) 
						RETURNING ID;`

	// Execute the query and get the last inserted ID
	err := r.db.QueryRow(
		query,
		book.Title,
		book.Author,
		book.Year,
	).Scan(&book.ID)

	// Check if there is an error
	if err != nil {
		return nil, err
	}

	// Return the new book
	return book, nil
}
