package main

import (
	"fmt"
	"log"
	"net/http"
)

type fuga int

func (f *fuga) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "fuga type: %d", *f)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func main() {

	http.Handle("/assets/", http.FileServer(http.Dir("./")))

	http.HandleFunc("/hello", hello)

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "world")
	})

	f := fuga(1)
	http.Handle("/", &f)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
