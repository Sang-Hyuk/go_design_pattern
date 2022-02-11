package main

import "fmt"

type Screen interface {
    Up()
    Down()
}

type WideScreen struct {

}

func (w WideScreen) Up() {
	fmt.Println("WideScreen Up.....")
}

func (w WideScreen) Down() {
	fmt.Println("WideScreen Down.....")
}

type Power interface {
    TurnOn()
    TurnOff()
}

type Projector struct {

}

func (p Projector) TurnOn() {
	fmt.Println("Turn On Projector.....")
}

func (p Projector) TurnOff() {
	fmt.Println("Turn Off Projector.....")
}

type Amp struct {

}

func (a Amp) TurnOn() {
	fmt.Println("Turn On Amp.....")
}

func (a Amp) TurnOff() {
	fmt.Println("Turn On Amp.....")
}

type HomeTheaterFacade interface {
    WatchMovie()
    EndMovie()
}

type HomeTheater struct {
    amp Amp
    projector Projector
    wideScreen WideScreen
}

func NewHomeTheater(amp Amp, projector Projector, wideScreen WideScreen) *HomeTheater {
	return &HomeTheater{amp: amp, projector: projector, wideScreen: wideScreen}
}

func (h HomeTheater) WatchMovie() {
	h.wideScreen.Down()
	h.projector.TurnOn()
	h.amp.TurnOn()
}

func (h HomeTheater) EndMovie() {
	h.amp.TurnOff()
	h.projector.TurnOff()
	h.wideScreen.Up()
}

func main() {
	homeTheater := NewHomeTheater(Amp{}, Projector{}, WideScreen{})
	homeTheater.WatchMovie()
	homeTheater.EndMovie()
}
