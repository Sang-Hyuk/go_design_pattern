package main

import "fmt"

type Subject interface {
    RegisterObserver(o Observer)
    RemoveObserver(o Observer)
    NotifyObservers()
}

type Observer interface {
    Update(temp float64, humidity float64, pressure float64)
}

type DisplayElement interface {
    Display()
}

type WeatherData struct {
	observers []Observer
	temperature float64
	humidity float64
	pressure float64
}

func (w *WeatherData) RegisterObserver(o Observer) {
	if o == nil  {
		return
	}
	w.observers = append(w.observers, o)
}

func (w *WeatherData) RemoveObserver(o Observer) {
	if o == nil  {
		return
	}
	for i, observer := range w.observers {
		fmt.Printf("[%p] [%p]\n", observer, o)
		if o == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			return
		}
	}
}

func (w WeatherData) NotifyObservers() {
	for _, observer := range w.observers {
		if observer == nil {
			continue
		}
		observer.Update(w.temperature, w.humidity, w.pressure)
	}
}

func (w WeatherData) MeasurementsChanged()  {
	w.NotifyObservers()
}

func (w *WeatherData) SetMeasurements(temperature float64, humidity float64, pressure float64) {
	w.temperature = temperature
	w.humidity = humidity
	w.pressure = pressure
	w.MeasurementsChanged()
}

type CurrentConditionsDisplay struct {
    temperature float64
    humidity float64
	pressure float64
    weatherData Subject
}

func NewCurrentConditionsDisplay(weatherData Subject) *CurrentConditionsDisplay {
	c := &CurrentConditionsDisplay{weatherData: weatherData}
	weatherData.RegisterObserver(c)
	return c
}

func (c *CurrentConditionsDisplay) Update(temp float64, humidity float64, pressure float64) {
	c.temperature = temp
	c.humidity = humidity
	c.pressure = pressure
	c.Display()
}

func (c CurrentConditionsDisplay) Display() {
	fmt.Printf("현재 상태 : 기온[%f] 습도[%f] 기압[%f]\n", c.temperature, c.humidity, c.pressure)
}

type StatisticsDisplay struct {
    temperatures []float64
    weatherData Subject
}

func NewStatisticsDisplay(weatherData Subject) *StatisticsDisplay {
	c := &StatisticsDisplay{weatherData: weatherData}
	weatherData.RegisterObserver(c)
	return c
}

func (s *StatisticsDisplay) Update(temp float64, humidity float64, pressure float64) {
	s.temperatures = append(s.temperatures, temp)
	s.Display()
}

func (s StatisticsDisplay) Display() {
	var avg, sum, min, max float64
	for _, temperature := range s.temperatures {
		if min == 0 || min > temperature {
			min = temperature
		}

		if max == 0 || max < temperature{
			max = temperature
		}

		sum += temperature
	}

	avg = sum / float64(len(s.temperatures))

	fmt.Printf("평균/최대/최소 온도 = %f,%f,%f\n",avg,max,min)
}

func main() {
	weatherData := &WeatherData{}
	c := NewCurrentConditionsDisplay(weatherData)
	d := NewStatisticsDisplay(weatherData)

	weatherData.SetMeasurements(80, 65, 30.4)
	weatherData.SetMeasurements(82, 70, 29.2)
	weatherData.SetMeasurements(78, 90, 20.5)

	weatherData.RemoveObserver(d)

	weatherData.SetMeasurements(75, 91, 22.5)

	weatherData.RemoveObserver(c)

	weatherData.SetMeasurements(79, 93, 24.5)
}
