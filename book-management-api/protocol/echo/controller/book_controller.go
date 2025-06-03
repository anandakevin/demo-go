package controller

import (
	"book-management-api/domain/dto"
	"book-management-api/domain/entity"
	"book-management-api/domain/usecase"
	"book-management-api/internal/parser"
	"book-management-api/protocol/echo/response"
	echo_validator "book-management-api/protocol/echo/validator"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookController struct {
	usecase usecase.IBookUsecase
}

func NewBookController(
	bookUsecase usecase.IBookUsecase,
) *BookController {
	return &BookController{
		usecase: bookUsecase,
	}
}

func (c *BookController) GetBooks(ctx echo.Context) error {
	// Initialize pagination with defaults
	pagination := dto.PaginationRequest{}
	if err := echo_validator.Bind(ctx, &pagination); err != nil {
		return response.Error(ctx, http.StatusBadRequest, errors.New("Invalid pagination parameters"))
	}

	result, err := c.usecase.GetBooks(pagination)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	return response.Success(ctx, http.StatusOK, result)
}

func (c *BookController) GetBookByISBN(ctx echo.Context) error {
	var params dto.ISBNParam
	if err := echo_validator.Bind(ctx, &params); err != nil {
		return response.Error(ctx, http.StatusBadRequest, errors.New("Invalid isbn"))
	}

	result, err := c.usecase.GetBookByISBN(params.ISBN)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	return response.Success(ctx, http.StatusOK, result)
}

func (c *BookController) CreateBook(ctx echo.Context) error {
	// Get the book from the request body
	var bookDto dto.CreateBook
	if err := echo_validator.Bind(ctx, &bookDto); err != nil {
		return response.Error(ctx, http.StatusBadRequest, errors.New("Invalid request"))
	}

	releaseDate, err := parser.ParseDate(bookDto.ReleaseDate)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	bookEntity := entity.Book{
		Title:       bookDto.Title,
		Author:      bookDto.Author,
		ISBN:        bookDto.ISBN,
		ReleaseDate: releaseDate,
	}

	result, err := c.usecase.CreateBook(bookEntity)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	resultResponse := dto.BookResponse{
		Title:       result.Title,
		Author:      result.Author,
		ISBN:        result.ISBN,
		ReleaseDate: result.ReleaseDate,
	}

	return response.Success(ctx, http.StatusCreated, resultResponse)
}

func (c *BookController) UpdateBookByISBN(ctx echo.Context) error {
	// Get the book from the request body
	var bookDto dto.UpdateBook
	if err := echo_validator.Bind(ctx, &bookDto); err != nil {
		return response.Error(ctx, http.StatusBadRequest, errors.New("Invalid request"))
	}

	releaseDate, err := parser.ParseDate(bookDto.ReleaseDate)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	bookEntity := entity.Book{
		Title:       bookDto.Title,
		Author:      bookDto.Author,
		ISBN:        bookDto.ISBN,
		ReleaseDate: releaseDate,
	}

	result, err := c.usecase.UpdateBook(bookEntity)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	resultResponse := dto.BookResponse{
		Title:       result.Title,
		Author:      result.Author,
		ISBN:        result.ISBN,
		ReleaseDate: result.ReleaseDate,
	}

	return response.Success(ctx, http.StatusOK, resultResponse)
}

func (c *BookController) DeleteBookByISBN(ctx echo.Context) error {
	// Get the book by isbn in the path variable
	var params dto.ISBNParam
	if err := echo_validator.Bind(ctx, &params); err != nil {
		return response.Error(ctx, http.StatusBadRequest, errors.New("Invalid isbn"))
	}

	result, err := c.usecase.DeleteBookByISBN(params.ISBN)
	if err != nil {
		return response.Error(ctx, http.StatusBadRequest, err)
	}

	resultResponse := dto.BookResponse{
		Title:       result.Title,
		Author:      result.Author,
		ISBN:        result.ISBN,
		ReleaseDate: result.ReleaseDate,
	}

	return response.Success(ctx, http.StatusOK, resultResponse)
}
