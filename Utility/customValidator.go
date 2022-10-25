package Utility

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	validate := validator.New()
	errvalidate := validate.Struct(i)
	if errvalidate != nil {
		if strings.Contains(errvalidate.Error(), "FirstName") {
			return echo.NewHTTPError(http.StatusBadRequest, "firstname is not valid")
		} else if strings.Contains(errvalidate.Error(), "LastName") {
			return echo.NewHTTPError(http.StatusBadRequest, "LastName is not valid")
		} else if strings.Contains(errvalidate.Error(), "BirthLocation") {
			return echo.NewHTTPError(http.StatusBadRequest, "BirthLocation is not 28 countries of iran")
		} else if strings.Contains(errvalidate.Error(), "BirthDate") {
			return echo.NewHTTPError(http.StatusBadRequest, "BirthDate is not a valid birth number")
		} else if strings.Contains(errvalidate.Error(), "PhoneNumber") {
			return echo.NewHTTPError(http.StatusBadRequest, "PhoneNumber is not valid check if your PhoneNumber start with 09 ")
		} else if strings.Contains(errvalidate.Error(), "Email") {
			return echo.NewHTTPError(http.StatusBadRequest, "Email is not a valid format")
		}

		//return echo.NewHTTPError(http.StatusBadRequest, errvalidate.Error())
	}

	return nil
}
