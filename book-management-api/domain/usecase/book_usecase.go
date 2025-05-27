package usecase

import (
	"book-management-api/domain/dto"
	"book-management-api/domain/entity"
	"book-management-api/internal/logger"
	"errors"
	"fmt"
	"sort"
)

type bookUsecase struct {
	logger logger.Logger
}

type IBookUsecase interface {
	GetBooks(pagination dto.PaginationRequest) (dto.PaginatedResponse[entity.Book], error)
	GetBookByISBN(isbn string) (*entity.Book, error)
	CreateBook(book entity.Book) (*entity.Book, error)
	UpdateBook(book entity.Book) (*entity.Book, error)
	DeleteBookByISBN(isbn string) (*entity.Book, error)
}

// NewBankService new bank service
func NewBookUsecase(logger logger.Logger) *bookUsecase {
	return &bookUsecase{
		logger: logger,
	}
}

// Global store instance
var store = &entity.BookStore{
	Books: make(map[string]entity.Book),
}

// GetBooks handles retrieving all books with pagination
func (u *bookUsecase) GetBooks(pagination dto.PaginationRequest) (dto.PaginatedResponse[entity.Book], error) {
	store.Mutex.RLock()
	defer store.Mutex.RUnlock()

	// Convert map to slice for pagination
	var books []entity.Book
	for _, book := range store.Books {
		books = append(books, book)
	}

	// Sort books based on selected column
	sort.Slice(books, func(i, j int) bool {
		// Define comparison based on selected column
		var comparison bool
		switch pagination.SortBy {
		case "title":
			comparison = books[i].Title < books[j].Title
		case "author":
			comparison = books[i].Author < books[j].Author
		case "isbn":
			comparison = books[i].ISBN < books[j].ISBN
		case "release_date":
			comparison = books[i].ReleaseDate.Before(books[j].ReleaseDate)
		default:
			comparison = books[i].Title < books[j].Title
		}

		// Reverse comparison for descending order
		if pagination.SortOrder == "desc" {
			return !comparison
		}
		return comparison
	})

	// pagination
	paginatedBooks, total := PaginateBooks(books, pagination.Page, pagination.Limit)
	totalPages := (total + pagination.Limit - 1) / pagination.Limit

	paginatedResponse := dto.PaginatedResponse[entity.Book]{
		Page:       pagination.Page,
		Limit:      pagination.Limit,
		TotalPages: totalPages,
		Total:      total,
		Data:       paginatedBooks,
	}

	return paginatedResponse, nil
}

func PaginateBooks(books []entity.Book, page, limit int) ([]entity.Book, int) {
	total := len(books)
	start := (page - 1) * limit
	end := start + limit

	if start >= total {
		return []entity.Book{}, total
	}

	if end > total {
		end = total
	}

	return books[start:end], total
}

// GetBookByISBN handles retrieving a single book by ISBN
func (u *bookUsecase) GetBookByISBN(isbn string) (*entity.Book, error) {
	store.Mutex.RLock()
	defer store.Mutex.RUnlock()

	book, exists := store.Books[isbn]
	if !exists {
		return nil, errors.New("Book not found")
	}

	return &book, nil
}

// CreateBook handles book creation
func (u *bookUsecase) CreateBook(book entity.Book) (*entity.Book, error) {

	store.Mutex.Lock()
	defer store.Mutex.Unlock()

	// Check if ISBN already exists
	if _, exists := store.Books[book.ISBN]; exists {
		return nil, errors.New("Book already exists")
	}

	// Store the book
	store.Books[book.ISBN] = book

	// Log asynchronously
	u.logger.Info(fmt.Sprintf("Book created: %s", book.ISBN))

	return &book, nil
}

// UpdateBookByISBN handles updating a book
func (u *bookUsecase) UpdateBook(book entity.Book) (*entity.Book, error) {

	store.Mutex.Lock()
	defer store.Mutex.Unlock()

	// Check if book exists
	if _, exists := store.Books[book.ISBN]; !exists {
		return nil, errors.New("Book not found")
	}

	// Update the book
	store.Books[book.ISBN] = book

	// Log asynchronously
	u.logger.Info(fmt.Sprintf("Book updated: %s", book.ISBN))

	return &book, nil
}

// DeleteBookByISBN handles deleting a book
func (u *bookUsecase) DeleteBookByISBN(isbn string) (*entity.Book, error) {
	store.Mutex.Lock()
	defer store.Mutex.Unlock()

	// Check if book exists
	if _, exists := store.Books[isbn]; !exists {
		return nil, errors.New("Book not found")
	}

	// Get the book before deleting
	book := store.Books[isbn]

	// Delete the book
	delete(store.Books, isbn)

	// Log asynchronously
	u.logger.Info(fmt.Sprintf("Book deleted: %s", isbn))

	return &book, nil
}
