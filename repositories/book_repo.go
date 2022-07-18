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

func (r bookRepository) GetBooks() ([]models.Book, error) {
	
	return []models.Book{}, nil
}

// CreateBook creates a new book
func (r bookRepository) CreateBook(book *models.Book) (*models.Book, error) {
	// Raw SQL query
	query := `INSERT INTO books 
						(title, author, year) 
						VALUES ($1, $2, $3) 
						RETURNING ID;`

	// Begin transaction
	tx, err := r.db.Begin()
	// Check error
	if err != nil {
		return nil, err
	}

	// Prepare query
	prepare, err := tx.Prepare(query)
	// Check error
	if err != nil {
		return nil, err
	}

	// Execute the query and get the last inserted ID
	err = tx.Stmt(prepare).QueryRow(
		query,
		book.Title,
		book.Author,
		book.Year,
	).Scan(&book.ID)

	// Check if there is an error
	if err != nil {
		// Rollback the transaction
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	tx.Commit()

	// Return the new book
	return book, nil
}
