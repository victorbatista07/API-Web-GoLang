package routers

import (
	"github.com/go-chi/chi"
	"github.com/webApplication/controllers"
)

// init gives to App a new value of echo
func init () {
	App := chi.NewRouter()
	App.Get("/", controllers.Hello)
}

