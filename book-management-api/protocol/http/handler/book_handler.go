package handler

import (
	"book-management-api/domain/dto"
	"book-management-api/domain/entity"
	"book-management-api/domain/usecase"
	"book-management-api/internal/parser"
	"book-management-api/protocol/http/response"
	"book-management-api/protocol/http/validator"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	usecase usecase.IBookUsecase
}

func NewBookHandler(bookUsecase usecase.IBookUsecase) *BookHandler {
	return &BookHandler{
		usecase: bookUsecase,
	}
}

// GetBooksHandler handles GET /books with pagination
func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	// Parse pagination parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortBy := r.URL.Query().Get("sort_by")
	sortOrder := r.URL.Query().Get("sort_order")

	page := 1
	limit := 10

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 100 {
			limit = l
		}
	}

	paginationReq := dto.PaginationRequest{
		Page:      page,
		Limit:     limit,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}

	paginatedResponse, err := h.usecase.GetBooks(paginationReq)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.SendJSONResponse(w, paginatedResponse, http.StatusOK)
}

// GetBookByISBNHandler handles GET /books/{isbn}
func (h *BookHandler) GetBookByISBN(w http.ResponseWriter, r *http.Request) {
	isbn := h.extractISBNFromPath(r.URL.Path)
	if isbn == "" {
		response.SendErrorResponse(w, "ISBN is required", http.StatusBadRequest)
		return
	}

	book, err := h.usecase.GetBookByISBN(isbn)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	response.SendJSONResponse(w, book, http.StatusOK)
}

// CreateBookHandler handles POST /books
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookDto dto.CreateBook
	if err := json.NewDecoder(r.Body).Decode(&bookDto); err != nil {
		response.SendErrorResponse(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	releaseDate, err := parser.ParseDate(bookDto.ReleaseDate)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	bookEntity := entity.Book{
		Title:       bookDto.Title,
		Author:      bookDto.Author,
		ISBN:        bookDto.ISBN,
		ReleaseDate: releaseDate,
	}

	if err = validator.ValidateBook(bookEntity); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdBook, err := h.usecase.CreateBook(bookEntity)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusConflict)
		return
	}

	response.SendJSONResponse(w, createdBook, http.StatusCreated)
}

// UpdateBookHandler handles PUT /books/{isbn}
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	isbn := h.extractISBNFromPath(r.URL.Path)
	if isbn == "" {
		response.SendErrorResponse(w, "ISBN is required", http.StatusBadRequest)
		return
	}

	var bookDto dto.UpdateBook
	if err := json.NewDecoder(r.Body).Decode(&bookDto); err != nil {
		response.SendErrorResponse(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	releaseDate, err := parser.ParseDate(bookDto.ReleaseDate)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
	}

	bookEntity := entity.Book{
		Title:       bookDto.Title,
		Author:      bookDto.Author,
		ISBN:        isbn,
		ReleaseDate: releaseDate,
	}

	if err = validator.ValidateBook(bookEntity); err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedBook, err := h.usecase.UpdateBook(bookEntity)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			response.SendErrorResponse(w, err.Error(), http.StatusNotFound)
		} else {
			response.SendErrorResponse(w, err.Error(), http.StatusConflict)
		}
		return
	}

	response.SendJSONResponse(w, updatedBook, http.StatusOK)
}

// DeleteBookHandler handles DELETE /books/{isbn}
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	isbn := h.extractISBNFromPath(r.URL.Path)
	if isbn == "" {
		response.SendErrorResponse(w, "ISBN is required", http.StatusBadRequest)
		return
	}

	book, err := h.usecase.DeleteBookByISBN(isbn)
	if err != nil {
		response.SendErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	response.SendJSONResponse(w, book, http.StatusOK)
}

// Helper method to extract ISBN from URL path
func (h *BookHandler) extractISBNFromPath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) >= 3 && parts[1] == "books" {
		return parts[2]
	}
	return ""
}
