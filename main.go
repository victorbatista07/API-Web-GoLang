package main

import (
	"github.com/webApplication/routers"
	"github.com/labstack/echo/middleware"
)

func main () {
	e := routers.App
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":3000"))

}
