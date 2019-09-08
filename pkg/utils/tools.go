package utils

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io/ioutil"
)

func GetRoutes(e *echo.Echo) {
	data, _ := json.MarshalIndent(e.Routes(), "", "  ")
	ioutil.WriteFile("routes.json", data, 0644)
}
