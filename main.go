package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"example/Weather-App/weather"
)

func main() {

	r := gin.Default()

	// Define your route for weather forecast
	r.GET("/weather", weather.WeatherForecastHandler)

	port := 8080 // Choose the port you want to use

	fmt.Printf("Server is running on :%d...\n", port)
	r.Run(fmt.Sprintf(":%d", port))
}
