package controller

import (
	"log"
)

type DemoController struct{}

func (d DemoController) GreenOn() {
	log.Println("Green on.")
}

func (d DemoController) YellowOn() {
	log.Println("Yellow on.")
}

func (d DemoController) RedOn() {
	log.Println("Red on.")
}

func (d DemoController) YellowRedOn() {
	log.Println("Yellow-Red on.")
}

func (d DemoController) AllOn() {
	log.Println("All on.")
}

func (d DemoController) GreenOff() {
	log.Println("Green off.")
}

func (d DemoController) YellowOff() {
	log.Println("Yellow off.")
}

func (d DemoController) RedOff() {
	log.Println("Red off.")
}

func (d DemoController) YellowRedOff() {
	log.Println("Yellow-Red off.")
}

func (d DemoController) AllOff() {
	log.Println("All off.")
}
