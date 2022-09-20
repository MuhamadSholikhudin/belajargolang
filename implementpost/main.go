package main

import (
	"belajargolang/implementpost/views"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	homeTpl    *views.View
	aboutusTpl *views.View
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTpl.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}
func aboutusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := aboutusTpl.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}
func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Halaman yang dicari tidak ditemukan</h1>")
}
func main() {
	homeTpl = views.NewView("views/home.gohtml")
	aboutusTpl = views.NewView("views/aboutus.gohtml")
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notfoundHandler)
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/aboutus", aboutusHandler)
	http.ListenAndServe(":3000", r)
}
