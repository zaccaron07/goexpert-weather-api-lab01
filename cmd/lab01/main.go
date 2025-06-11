package main

import (
	webserver "github.com/zaccaron07/goexpert-weather-api-lab01/internal/infra/web/webserver"
)

func main() {
	webserver := webserver.NewWebServer("127.0.0.1:8080")
	webWeatherHandler := NewWeatherHandler()
	webserver.AddHandler("/zipcode/{zipcode}/weather", webWeatherHandler.Fetch)

	webserver.Start()
}
