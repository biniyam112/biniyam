package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		remover(w, r)
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles(("src/github.com/snippetboxx/ui/html/home.page.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "server error that is internal", 500)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "server error(That is internal)", 500)
	}
}

func remover(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(405)
		http.Error(w, "Unsupported request", 450)
		return
	}
	w.Write([]byte(`{"name":"Bini"}` + "\n"))
}

func another(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	// w.Header()["Date"] = nil
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "This is the collected id %d:", id)

	fmt.Print("")
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Unsupported request", 450)
		return
	}
	w.Write([]byte(`{"name":"Alex"}` + "\n"))
}
