package controllers

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Custumer struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"gte=17,lte=35"`
}

func TestVariabelValidation(c echo.Context) error {
	v := validator.New()

	email := "bams@gmail.com"
	err := v.Var(email, "required,email")
	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Custumer{
		Nama:   "Bass",
		Email:  "not@gmail.com",
		Alamat: "Kudus",
		Umur:   18,
	}

	err := v.Struct(cust)
	if err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}
