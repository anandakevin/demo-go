package response

import (
	"book-management-api/domain/dto"

	"github.com/labstack/echo/v4"
)

func Success(ctx echo.Context, status int, data interface{}) error {
	return ctx.JSON(status, dto.SuccessResponse{
		Status: "success",
		Data:   data,
	})
}

func Error(ctx echo.Context, status int, err error) error {
	return ctx.JSON(status, dto.ErrorResponse{
		Status: "error",
		Error:  err.Error(),
	})
}
