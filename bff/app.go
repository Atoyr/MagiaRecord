package main

import "github.com/atoyr/MagiaRecord/bff/handler"

func main() {
	e := echo.New()

	e.GET(handler.GetMagicalGirls)
}
