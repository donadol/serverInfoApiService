package main

import (
	"net/http"
	"time"

	"./controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	//r.Get("/servers", controllers.InfoServers)
	r.Get("/{domain}", controllers.InfoDomain)

	http.ListenAndServe(":3000", r)
}