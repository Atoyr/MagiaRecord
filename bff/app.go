package main

import (
	"github.com/atoyr/MagiaRecord/bff/handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			con := handler.Context{Context: c, DriverName: "", DataSourceName: ""}
			return h(&con)
		}
	})
	e.GET("/magical_girls", c(handler.GetMagicalGirlsHandler()))
}

func c(h handler.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*handler.Context))
	}
}
