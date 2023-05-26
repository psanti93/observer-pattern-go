package main

import (
	display "github.com/psanti93/observer-pattern-go/displays"
	weatherData "github.com/psanti93/observer-pattern-go/weatherdata"
)

func main() {

	weatherData := &weatherData.WeatherData{}

	currentDisplay := &display.CurrentDisplay{
		WeatherData: weatherData,
	}

	forecastDisplay := &display.Forecast{
		WeatherData: weatherData,
	}

	weatherData.RegisterObserver(currentDisplay)
	weatherData.RegisterObserver(forecastDisplay)

	weatherData.SetMeasurements(100, 90, 80)
	weatherData.SetMeasurements(60, 65, 100)
	weatherData.SetMeasurements(45, 40, 100)

}
