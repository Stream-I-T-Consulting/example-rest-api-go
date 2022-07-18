package repositories

import (
	"database/sql"
	"errors"
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
	// Raw SQL query
	query := `SELECT 
						id, title, author, year 
						FROM books;`

	// Get the books
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	// Close the rows database connection
	defer rows.Close()

	// Create a slice to store the books
	var books []models.Book

	// Loop through the rows
	for rows.Next() {
		// Create a new book
		var book models.Book

		// Get the columns from the row
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		if err != nil {
			return nil, err
		}

		// Append the book to the slice
		books = append(books, book)
	}

	// Prepare query
	return books, nil
}

func (r bookRepository) GetBookByID(id int) (*models.Book, error) {
	// Raw SQL query
	query := `SELECT 
						id, title, author, year 
						FROM books
						WHERE id = $1
						LIMIT 1;`

	// Prepare query
	stmt, err := r.db.Prepare(query)
	// Check error
	if err != nil {
		return nil, err
	}

	// Get the book
	var book models.Book

	// Execute the query and get the book
	err = stmt.QueryRow(id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.Year,
	)
	// Check error
	if err != nil {
		return nil, errors.New("book not found")
	}

	// Close the database connection
	defer stmt.Close()

	// Prepare query
	return &book, nil
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
