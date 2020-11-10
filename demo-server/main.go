package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
)

type Response struct {
	Message string `json:"message" xml:"message"`
	Status  string `json:"status" xml:"status"`
}

// Handler
func (r Response) HelloHandlerPrettyJson(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, r, "  ")
}

func (r Response) HelloHandlerPrettyXML(c echo.Context) error {
	return c.XMLPretty(http.StatusOK, r, "  ")
}

func (r Response) HelloHandler(c echo.Context) error {
	response := fmt.Sprintf("<h1>%s</h1>", r.Message)
	return c.HTML(http.StatusOK, response)
}

func (r Response) HelloHandlerYaml(c echo.Context) error {
	MIMEApplicationYAML := "application/vnd.yaml;charset=UTF-8"
	c.Response().Header().Set(echo.HeaderContentType, MIMEApplicationYAML)
	c.Response().WriteHeader(http.StatusOK)
	return yaml.NewEncoder(c.Response()).Encode(r)
}

func main() {
	e := echo.New()
	r := Response{
		Message: "Hello from Echo!",
		Status:  "OK",
	}
	// make routes
	e.GET("/", r.HelloHandler)
	e.GET("/json", r.HelloHandlerPrettyJson)
	e.GET("/xml", r.HelloHandlerPrettyXML)
	e.GET("/yaml", r.HelloHandlerYaml)
	e.Logger.Fatal(e.Start(":1323"))
}
