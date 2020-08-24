package lights

import (
	"log"
)

// NegatedTrafficLight operates on the GPIO pins with a negated high low configuration.
type NegatedTrafficLight struct {
	gpioController ClassicTrafficLight
}

// NewNegatedTrafficLight creates a new traffic light configuration.
func NewNegatedTrafficLight(greenPin int, yellowPin int, redPin int) *NegatedTrafficLight {
	n := new(NegatedTrafficLight)
	n.gpioController = *NewClassicTrafficLight(greenPin, yellowPin, redPin)
	return n
}

// GreenOn turns the green light on.
func (n *NegatedTrafficLight) GreenOn() {
	log.Println("Green on. (negated call)")
	n.gpioController.GreenOff()
}

// YellowOn turns the yellow light on.
func (n *NegatedTrafficLight) YellowOn() {
	log.Println("Yellow on. (negated call)")
	n.gpioController.YellowOff()
}

// RedOn turns the red light on.
func (n *NegatedTrafficLight) RedOn() {
	log.Println("Red on. (negated call)")
	n.gpioController.RedOff()
}

// YellowRedOn turns the yellow and red light on.
func (n *NegatedTrafficLight) YellowRedOn() {
	log.Println("Yellow-Red on. (negated call)")
	n.gpioController.YellowRedOff()
}

// AllOn turns the green, yellow and red lights on.
func (n *NegatedTrafficLight) AllOn() {
	log.Println("All on. (negated call)")
	n.gpioController.AllOff()
}

// GreenOff turns the green light off.
func (n *NegatedTrafficLight) GreenOff() {
	log.Println("Green off. (negated call)")
	n.gpioController.GreenOn()
}

// YellowOff turns the yellow light off.
func (n *NegatedTrafficLight) YellowOff() {
	log.Println("Yellow off. (negated call)")
	n.gpioController.YellowOn()
}

// RedOff turns the red light off.
func (n *NegatedTrafficLight) RedOff() {
	log.Println("Red off. (negated call)")
	n.gpioController.RedOn()
}

// YellowRedOff turns the yellow and red light off.
func (n *NegatedTrafficLight) YellowRedOff() {
	log.Println("Yellow-Red off. (negated call)")
	n.gpioController.YellowRedOn()
}

// AllOff turns the green, yellow and red lights off.
func (n *NegatedTrafficLight) AllOff() {
	log.Println("All off. (negated call)")
	n.gpioController.AllOn()
}
