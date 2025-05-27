package validator

import (
	"net/http"

	internal_validator "book-management-api/internal/validator"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// EchoValidator implements internal_validator.RequestValidator
type EchoValidator struct {
	Validator *validator.Validate
}

func NewEchoValidator() *EchoValidator {
	return &EchoValidator{
		Validator: validator.New(),
	}
}

func (v *EchoValidator) Validate(i interface{}) error {
	if err := v.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Bind(ctx echo.Context, i interface{}) error {
	if err := ctx.Bind(i); err != nil {
		return err
	}
	internal_validator.SetDefaults(i)
	return ctx.Validate(i)
}
