package controllers

import (
	"fmt"
	"net/http"
	"time"

	"belajargolang/ECHO-REST/helpers"
	"belajargolang/ECHO-REST/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin string `json:"admin"`
	jwt.StandardClaims
}

func GenerateHashPAssword(c echo.Context) error {
	password := c.Param("password")

	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := models.CheckLogin(username, password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}
	/*
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Berhasil Login",
		})

	*/

	// token := jwt.New(jwt.SigningMethodES256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["name"] = username
	// claims["level"] = "application"
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Set custom claims
	claims := &jwtCustomClaims{
		username,
		"application",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("message key is of invalid type")

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}
