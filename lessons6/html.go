package main

import (
	"net/http"
)

func main() {
	// func(w http.ResponseWriter, r *http.Request)
	http.DefaultServeMux.HandleFunc("/privet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`privet`))
	})
	http.DefaultServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello!`))
	})
	http.DefaultServeMux.HandleFunc("/hello&name=World", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`key = name, value = World`))
	})
	server := &http.Server{
		Addr:    ":3333",
		Handler: nil,
	}
	server.ListenAndServe()

}
