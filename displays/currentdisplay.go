package displays

import (
	"fmt"

	weatherData "github.com/psanti93/observer-pattern-go/weatherdata"
)

type CurrentDisplay struct {
	WeatherData *weatherData.WeatherData
}

func (c *CurrentDisplay) Display() {
	fmt.Printf("The Current forecast for today is %d degrees with %d humidity \n", c.WeatherData.Temperature, c.WeatherData.Humidity)
}

func (c *CurrentDisplay) Update() {
	c.Display()
}
