package main

import (
	"httpserver/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.Home)
	mux.HandleFunc("/animal", handler.GetAnimal)
	mux.HandleFunc("/animal/:id", handler.GetAnimal)

	err := http.ListenAndServe(
		":4500", mux,
	)
	if err != nil {
		panic(err)
	}
}
