package interfaces

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
