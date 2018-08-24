package controllers

import (
	"net/http"
)

// Hello it's API's homepage
func Hello (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá fucking shit!"))
}
