package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

)

func fetchWeatherData(apiKey, location string) (*WeatherResponse, error) {

	//* Call Weather API
	url := fmt.Sprintf("https://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=no&alerts=no", apiKey, location)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//* Read Response Body
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}


	//* Parse responseBody Bytes to Strust Models
	var weather WeatherResponse
	if err := json.Unmarshal(data, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}
