package displays

import (
	"fmt"

	"github.com/psanti93/observer-pattern-go/weatherdata"
)

type Forecast struct {
	WeatherData *weatherdata.WeatherData
}

func (f *Forecast) Display() {
	if f.WeatherData.Temperature >= 90 && f.WeatherData.Humidity >= 90 {
		fmt.Println("Today's forecast calls for a very hot day!\n")
	} else if (f.WeatherData.Temperature <= 60 && f.WeatherData.Temperature >= 50) && (f.WeatherData.Humidity <= 70 && f.WeatherData.Humidity >= 60) {
		fmt.Println("Today's forecast calls for a nice cool day! \n")
	} else if f.WeatherData.Temperature <= 50 && f.WeatherData.Humidity <= 40 {
		fmt.Println("Expect a cold day \n")
	}
}

func (f *Forecast) Update() {
	f.Display()
}
