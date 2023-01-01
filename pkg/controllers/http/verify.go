package http

import (
	"github.com/go-playground/validator"
	"github.com/go-verify/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

type (
	Data struct {
		Name string `json:"name" validate:"required" `
	}

	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Verify(c echo.Context) error {
	c.Echo().Validator = &CustomValidator{validator: validator.New()}
	u := new(Data)
	err := c.Bind(u)
	if err = c.Validate(u); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	ofac := utils.Ofac(u.Name)

	return c.JSON(200, ofac)
}
