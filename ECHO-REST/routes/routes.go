package routes

import (
	"belajargolang/ECHO-REST/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is echo !")
	})

	r := e.Group("/api")

	r.GET("/pegawai", controllers.FetchAllPegawai)
	r.POST("/pegawai", controllers.StorePegawai)
	r.PUT("/pegawai", controllers.UpdatePegawai)
	r.DELETE("/pegawai", controllers.DeletePegawai)

	e.GET("/generate-hash/:password", controllers.GenerateHashPAssword)
	e.POST("/login", controllers.CheckLogin)

	e.GET("/test-struct-validation", controllers.TestStructValidation)
	e.GET("/test-variabel-validation", controllers.TestVariabelValidation)

	r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}))
	return e
}
