package v1

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/seyedmo30/Golang_URL_Shortener_Hexagonal/pkg/logger"
)

// /paths
func Routes(api *echo.Group) {

	expl := api.Group("/shortner")
	expl.POST("", Add)
	expl.GET("/:shortner", Get)

}
func Add(c echo.Context) error {

	ports := new(concritPorts)
	url := c.FormValue("url")

	shortner, err := ports.Add(context.Background(), url)
	if err != nil {
		return c.JSONBlob(500, []byte(`{"response": "error server","status":500}`))

	}

	return c.JSONBlob(200, []byte(`{"response":"`+shortner+`","status":200}`))
}

func Get(c echo.Context) error {
	ports := new(concritPorts)
	shortner := c.Param("shortner")
	logger.Log().Info(shortner)
	url, err := ports.Get(context.Background(), shortner)
	if err != nil {
		return c.JSONBlob(500, []byte(`{"response": "error server","status":500}`))

	}
	return c.JSONBlob(200, []byte(`{"response":"`+url+`","status":200}`))
}
