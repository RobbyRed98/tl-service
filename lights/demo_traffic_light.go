package lights

import (
	"log"
)

// DemoTrafficLight a demo traffic light which does not operate on the GPIO pins. Instead of that it logs the performed operations.
type DemoTrafficLight struct{}

// GreenOn logs that green light turned on.
func (d DemoTrafficLight) GreenOn() {
	log.Println("Green on.")
}

// YellowOn logs that yellow light turned on.
func (d DemoTrafficLight) YellowOn() {
	log.Println("Yellow on.")
}

// RedOn logs that red light turned on.
func (d DemoTrafficLight) RedOn() {
	log.Println("Red on.")
}

// YellowRedOn logs that yellow and red light turned on.
func (d DemoTrafficLight) YellowRedOn() {
	log.Println("Yellow-Red on.")
}

// AllOn logs that all lights turned on.
func (d DemoTrafficLight) AllOn() {
	log.Println("All on.")
}

// GreenOff logs that green light turned off.
func (d DemoTrafficLight) GreenOff() {
	log.Println("Green off.")
}

// YellowOff logs that yellow light turned off.
func (d DemoTrafficLight) YellowOff() {
	log.Println("Yellow off.")
}

// RedOff logs that red light turned off.
func (d DemoTrafficLight) RedOff() {
	log.Println("Red off.")
}

// YellowRedOff logs that yellow and red light turned off.
func (d DemoTrafficLight) YellowRedOff() {
	log.Println("Yellow-Red off.")
}

// AllOff logs that all lights turned off.
func (d DemoTrafficLight) AllOff() {
	log.Println("All off.")
}
