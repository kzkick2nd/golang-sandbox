package main

import (
	"fmt"
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
	http.HandleFunc("/hello", hello)

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "world")
	})

	f := fuga(1)
	http.Handle("/", &f)

	http.ListenAndServe(":8080", nil)
}
