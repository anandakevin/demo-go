package routes

import (
	"book-management-api/protocol/http/handler"
	"book-management-api/protocol/http/response"
	"net/http"
	"strings"
)

// BookRouter holds the handler dependencies
type BookRouter struct {
	bookHandler *handler.BookHandler
}

// NewBookRouter creates a new router with injected dependencies
func NewBookRouter(bookHandler *handler.BookHandler) *BookRouter {
	return &BookRouter{
		bookHandler: bookHandler,
	}
}

// Routes method handles routing logic
func (br *BookRouter) Routes(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch {
	case path == "/books" && r.Method == http.MethodPost:
		br.bookHandler.CreateBook(w, r)
	case path == "/books" && r.Method == http.MethodGet:
		br.bookHandler.GetBooks(w, r)
	case strings.HasPrefix(path, "/books/") && r.Method == http.MethodGet:
		br.bookHandler.GetBookByISBN(w, r)
	case strings.HasPrefix(path, "/books/") && r.Method == http.MethodPut:
		br.bookHandler.UpdateBook(w, r)
	case strings.HasPrefix(path, "/books/") && r.Method == http.MethodDelete:
		br.bookHandler.DeleteBook(w, r)
	default:
		response.SendErrorResponse(w, "Endpoint not found", http.StatusNotFound)
	}
}
