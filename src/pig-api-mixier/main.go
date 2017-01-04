package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"pig-api-mixier/client"
)

func latest(c echo.Context) error {
	pigs := client.Latest()
	b, _ := json.Marshal(pigs)
	return c.String(http.StatusOK, string(b))
}

func main() {
	fmt.Println("Start")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World of Pigs!!")
	})

	e.GET("/api/r/latest", latest)
	e.Logger.Debug(e.Start(":1323"))
}
