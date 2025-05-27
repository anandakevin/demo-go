package dto

type ISBNParam struct {
	ISBN string `param:"isbn" validate:"required,min=10,max=13,isbn"`
}

type CreateBook struct {
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Author      string `json:"author" validate:"required,min=1,max=100"`
	ISBN        string `json:"isbn" validate:"required,min=10,max=13,isbn"`
	ReleaseDate string `json:"release_date" validate:"required"`
}

type UpdateBook struct {
	Title       string `json:"title" validate:"required,min=1,max=200"`
	Author      string `json:"author" validate:"required,min=1,max=100"`
	ISBN        string `param:"isbn" validate:"required,min=10,max=13,isbn"`
	ReleaseDate string `json:"release_date" validate:"required"`
}
