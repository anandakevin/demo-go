package main

import (
	"book-management-api/domain/usecase"
	"book-management-api/internal/logger"
	"book-management-api/protocol/http/handler"
	"book-management-api/protocol/http/routes"
	"log"
	"net/http"
)

func main() {
	// Initialize dependencies in correct order

	// 1. Create Logger (lowest level dependency)
	loggerInstance := logger.NewAsyncLogger()
	defer loggerInstance.Close()

	// 2. Create Use Cases (business logic layer)
	bookUsecase := usecase.NewBookUsecase(loggerInstance)

	// 3. Create Handlers (presentation layer)
	bookHandler := handler.NewBookHandler(bookUsecase)

	// 4. Create Router with injected handler
	bookRouter := routes.NewBookRouter(bookHandler)

	// 5. Setup HTTP routes
	http.HandleFunc("/books", bookRouter.Routes)
	http.HandleFunc("/books/", bookRouter.Routes) // Handle paths with ISBN

	port := ":8080"
	loggerInstance.Info("Server is running on port " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
