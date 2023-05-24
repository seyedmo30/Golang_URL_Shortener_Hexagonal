package main

import (
	v1 "github.com/seyedmo30/Golang_URL_Shortener_Hexagonal/handler/http/echo/v1"
)

// @title          UrlShortner
// @version        1.0
// @description    UrlShortner

// @host     0.0.0.0:8000
// @BasePath /
// @schemes  http
func main() {

	port := "8008"
	host := "0.0.0.0"
	r := v1.New()
	//r.GET("/swagger/*", echoSwagger.WrapHandler)
	v1.Routes(r.Group("/api"))

	r.Logger.Fatal(r.Start(host + ":" + port))
}
