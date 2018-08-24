package main

import (
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/webApplication/controllers"
)

func main() {
    r := chi.NewRouter()
	r.Use(middleware.Logger)
    r.Get("/", controllers.Hello)
    http.ListenAndServe(":3000", r)
}


