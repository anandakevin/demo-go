package dto

type PaginationRequest struct {
	Page      int    `query:"page" default:"1" validate:"min=1"`
	Limit     int    `query:"limit" default:"10" validate:"min=1,max=100"`
	SortBy    string `query:"sort_by" default:"title" validate:"oneof=title author isbn release_date"`
	SortOrder string `query:"sort_order" default:"asc" validate:"oneof=asc desc"`
}
