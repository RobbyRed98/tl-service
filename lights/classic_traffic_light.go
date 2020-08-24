package lights

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

// ClassicTrafficLight operates on the GPIO pins with a regular high low configuration.
type ClassicTrafficLight struct {
	green  rpio.Pin
	yellow rpio.Pin
	red    rpio.Pin
}

// NewClassicTrafficLight creates a new classic traffic light.
func NewClassicTrafficLight(greenPin int, yellowPin int, redPin int) *ClassicTrafficLight {
	err := rpio.Open()
	if err != nil {
		log.Println("Cannot allocate RAM for GPIOs.")
		return nil
	}
	g := new(ClassicTrafficLight)
	g.green = rpio.Pin(greenPin)
	g.yellow = rpio.Pin(yellowPin)
	g.red = rpio.Pin(redPin)

	g.green.Output()
	g.yellow.Output()
	g.red.Output()
	return g
}

// GreenOn turns the green light on.
func (g *ClassicTrafficLight) GreenOn() {
	log.Println("Green on. (real call)")
	g.green.High()
}

// YellowOn turns the yellow light on.
func (g *ClassicTrafficLight) YellowOn() {
	log.Println("Yellow on. (real call)")
	g.yellow.High()
}

// RedOn turns the red light on.
func (g *ClassicTrafficLight) RedOn() {
	log.Println("Red on. (real call)")
	g.red.High()
}

// YellowRedOn turns the yellow and red light on.
func (g *ClassicTrafficLight) YellowRedOn() {
	log.Println("Yellow-Red on. (real call)")
	g.YellowOn()
	g.RedOn()
}

// AllOn turns the green, yellow and red lights on.
func (g *ClassicTrafficLight) AllOn() {
	log.Println("All on. (real call)")
	g.GreenOn()
	g.YellowRedOn()
}

// GreenOff turns the green light off.
func (g *ClassicTrafficLight) GreenOff() {
	log.Println("Green off. (real call)")
	g.green.Low()
}

// YellowOff turns the yellow light off.
func (g *ClassicTrafficLight) YellowOff() {
	log.Println("Yellow off. (real call)")
	g.yellow.Low()
}

// RedOff turns the red light off.
func (g *ClassicTrafficLight) RedOff() {
	log.Println("Red off. (real call)")
	g.red.Low()
}

// YellowRedOff turns the yellow and red light off.
func (g *ClassicTrafficLight) YellowRedOff() {
	log.Println("Yellow-Red off. (real call)")
	g.YellowOff()
	g.RedOff()
}

// AllOff turns the green, yellow and red lights off.
func (g *ClassicTrafficLight) AllOff() {
	log.Println("All off. (real call)")
	g.YellowRedOff()
	g.RedOff()
}
