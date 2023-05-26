package main

import "fmt"

type Subject interface {
	RegisterObserver(observer Observer)
	NotifyObservers()
}

type Observer interface {
	Update()
}

type Displace interface {
	Display()
}

type WeatherData struct {
	Observers   []Observer
	Temperature int
	Humidity    int
	Pressure    int
}

func (w *WeatherData) RegisterObserver(observer Observer) {
	w.Observers = append(w.Observers, observer)
}

func (w *WeatherData) NotifyObservers() {
	for _, observer := range w.Observers {
		observer.Update()
	}
}

type CurrentDisplay struct {
	WeatherData *WeatherData
}

func (c *CurrentDisplay) Display() {
	fmt.Printf("The Current forecast for today is %d degrees with %d humidity", c.WeatherData.Temperature, c.WeatherData.Humidity)
}

func (c *CurrentDisplay) Update() {
	c.Display()
}

func main() {

	weatherData := &WeatherData{
		Temperature: 60,
		Humidity:    30,
		Pressure:    20,
	}

	currentDisplay := &CurrentDisplay{
		WeatherData: weatherData,
	}

	weatherData.RegisterObserver(currentDisplay)
	weatherData.NotifyObservers()

}
