package routes

import (
	"book-management-api/protocol/echo/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BookRoutes(e *echo.Echo, ctrl *controller.BookController) {

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/books", ctrl.CreateBook)
	e.GET("/books", ctrl.GetBooks)
	e.GET("/books/:isbn", ctrl.GetBookByISBN)
	e.PUT("/books/:isbn", ctrl.UpdateBookByISBN)
	e.DELETE("/books/:isbn", ctrl.DeleteBookByISBN)
}
