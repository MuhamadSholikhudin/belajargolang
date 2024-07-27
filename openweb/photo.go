package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	http.HandleFunc("/photo", func(w http.ResponseWriter, r *http.Request) {

		u, err := url.Parse(r.RequestURI)
		if err != nil {
			log.Fatal(err)
		}
		q := u.Query()

		number_of_employee, checksearch := q["number_of_employee"]
		if checksearch != true {
			fmt.Println(err.Error())
			fmt.Println("query number of employee tidak ada")
		}

		fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%s.jpg", strings.Join(number_of_employee, "")))
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println("number of employee tidak ada foto")
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Write(fileBytes)
		return
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":9090", nil)
	fmt.Println("Listen on port 9090")
}

/*
package main

import (
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
	"net/url"
	"strings"

)

func main() {
	handler := http.HandlerFunc(HandleRequest)
	http.Handle("/photo", handler)
	http.ListenAndServe(":7070", nil)
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {


	u, err := url.Parse(r.RequestURI)
	if err != nil {
		log.Fatal(err)
	}
	q := u.Query()

	number_of_employee, checksearch := q["number_of_employee"]
	if checksearch != true {
		fmt.Println(err.Error())
	}

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%s.JPG", strings.Join(number_of_employee, "")))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}


*/
