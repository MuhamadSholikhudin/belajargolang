package main

import (
	"belajargolang/ECHO-REST/db"
	"belajargolang/ECHO-REST/routes"
)

func main() {

	db.Init()                        // Inisiasi Database
	e := routes.Init()               //Inisiasi Route
	e.Logger.Fatal(e.Start(":1234")) // Lisaten Port
}
