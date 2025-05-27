package dto

import (
	"time"
)

type BookResponse struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	ReleaseDate time.Time `json:"release_date"`
}
