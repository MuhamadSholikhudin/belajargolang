package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo! Apa kabar saya")

	})

	http.HandleFunc("/index", index)

	fmt.Println("starting web server at http://localhost:9000/")
	http.ListenAndServe(":9000", nil)
}
