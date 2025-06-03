package validator

import (
	"book-management-api/domain/entity"
	"fmt"
	"strings"
)

// Validation functions
func ValidateBook(book entity.Book) error {
	if strings.TrimSpace(book.Title) == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if strings.TrimSpace(book.Author) == "" {
		return fmt.Errorf("author cannot be empty")
	}
	if strings.TrimSpace(book.ISBN) == "" {
		return fmt.Errorf("ISBN cannot be empty")
	}
	return nil
}
