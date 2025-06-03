package entity

import (
	"sync"
	"time"
)

// Book represents a book entity
type Book struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	ReleaseDate time.Time `json:"release_date"`
}

// BookStore manages the in-memory storage of books
type BookStore struct {
	Books map[string]Book
	Mutex sync.RWMutex
}
