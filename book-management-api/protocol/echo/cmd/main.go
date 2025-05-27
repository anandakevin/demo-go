package main

import (
	"book-management-api/domain/usecase"
	"book-management-api/internal/logger"
	"book-management-api/protocol/echo/controller"
	"book-management-api/protocol/echo/routes"
	echo_validator "book-management-api/protocol/echo/validator"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	// Create Logger
	loggerInstance := logger.NewAsyncLogger()
	defer loggerInstance.Close() // Ensure logger is properly closed

	// Create Echo instance
	e := echo.New()
	e.Validator = &echo_validator.EchoValidator{Validator: validator.New()}

	// Usecases
	bookUsecase := usecase.NewBookUsecase(loggerInstance)

	// Controllers
	bookController := controller.NewBookController(bookUsecase)

	// Routes
	routes.BookRoutes(e, bookController)

	// Start server
	port := ":8080"
	loggerInstance.Info(fmt.Sprintf("Server starting on port %s\n", port))
	if err := e.Start(port); err != nil {
		loggerInstance.Error(fmt.Sprintf("Server failed to start: %v", err.Error()))
	}
}
