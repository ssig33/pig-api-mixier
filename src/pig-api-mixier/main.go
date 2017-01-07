package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"pig-api-mixier/client"
)

func latest(c echo.Context) error {
	pigs := client.Latest(c.QueryString())
	b, _ := json.Marshal(pigs)
	return c.String(http.StatusOK, string(b))
}

func search(c echo.Context) error {
	pigs := client.Search(c.QueryString())
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
	e.GET("/api/r/search", search)
	e.Logger.Debug(e.Start("0.0.0.0:" + os.Getenv("PORT")))
}
