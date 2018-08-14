package routers

import (
	"github.com/labstack/echo"
	"github.com/webApplication/controllers"
)

// App a variable type echo memory address
var App *echo.Echo

// init gives to App a new value of echo
func init () {
	App = echo.New()
	App.GET("/", controllers.Hello)
}

