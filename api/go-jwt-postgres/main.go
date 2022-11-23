package main

import (
	"fmt"
	"log"
	"net/http"

	"belajargolang/api/go-jwt-postgres/middlewares"

	"belajargolang/api/go-jwt-postgres/controllers/authcontroller"
	"belajargolang/api/go-jwt-postgres/controllers/productcontroller"
	"belajargolang/api/go-jwt-postgres/models"

	"github.com/gorilla/mux"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.Use(middlewares.JWTMiddleware)
	api.Use(middlewares.AccessMiddleware)
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")

	fmt.Println("Listen On PORT http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
