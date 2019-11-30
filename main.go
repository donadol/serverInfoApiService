package main

import (
	"fmt"
	"log"
	"net/http"

	"./models"
	"./controllers"

	_ "github.com/go-chi/chi"
)

const (
    dbName = "go-truora"
    dbHost = "localhost"
    dbPort = "33066"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", r)
}