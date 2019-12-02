package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", another)
	mux.HandleFunc("/snippet/create", remover)

	fmt.Println("starting servier")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
