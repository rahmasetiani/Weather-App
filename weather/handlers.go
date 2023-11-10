package weather

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

func WeatherForecastHandler(c *gin.Context) {

	//* Get Params
	location := c.Query("location")
	if location == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location is required"})
		return
	}

	weather, err := fetchWeatherData(API_KEY, location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching weather data"})
		return
	}

	//* Format response weather => responseTable
	responseTable := weather.Format()

	//* Return Response
	c.String(http.StatusOK, responseTable)
}
