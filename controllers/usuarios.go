package controllers

import (
	"net/http"
	"github.com/labstack/echo"
)

// Hello it's API's homepage
func Hello (context echo.Context) error {
	return context.String(http.StatusOK, "Olá")
}
