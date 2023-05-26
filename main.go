package main

import (
	currentDisplay "github.com/psanti93/observer-pattern-go/displays"
	weatherData "github.com/psanti93/observer-pattern-go/weatherdata"
)

func main() {

	weatherData := &weatherData.WeatherData{}

	currentDisplay := &currentDisplay.CurrentDisplay{
		WeatherData: weatherData,
	}

	weatherData.RegisterObserver(currentDisplay)
	weatherData.SetMeasurements(30, 60, 80)
	weatherData.SetMeasurements(20, 90, 100)

}
