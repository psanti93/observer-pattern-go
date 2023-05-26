package weatherdata

import interfaces "github.com/psanti93/observer-pattern-go/interfaces"

type WeatherData struct {
	Observers   []interfaces.Observer
	Temperature int
	Humidity    int
	Pressure    int
}

func (w *WeatherData) RegisterObserver(observer interfaces.Observer) {
	w.Observers = append(w.Observers, observer)
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.Observers {
		observer.Update()
	}
}

func (w *WeatherData) SetMeasurements(temperature int, humidity int, pressure int) {
	w.Temperature = temperature
	w.Humidity = humidity
	w.Pressure = pressure
	w.NotifyObservers()
}
