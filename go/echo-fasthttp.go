package main

import (
	"net/http"
	"fmt"

	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo"
)

const (
	port = ":3000"
)

func main() {
	e := echo.New()	
	
	fmt.Println(port)
	
	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "yo!")
	})
	
	e.Run(fasthttp.New(port))
}