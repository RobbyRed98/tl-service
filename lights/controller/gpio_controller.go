package controller

import (
	"log"

	"github.com/stianeikeland/go-rpio/v4"
)

type GpioController struct {
	green  rpio.Pin
	yellow rpio.Pin
	red    rpio.Pin
}

func NewGpioController(greenPinId int8, yellowPinId int8, redPinId int8) *GpioController {
	err := rpio.Open()
	if err != nil {
		log.Println("Cannot allocate RAM for GPIOs.")
		return nil
	}
	g := new(GpioController)
	g.green = rpio.Pin(greenPinId)
	g.yellow = rpio.Pin(yellowPinId)
	g.red = rpio.Pin(redPinId)

	g.green.Output()
	g.yellow.Output()
	g.red.Output()
	return g
}

func (g *GpioController) GreenOn() {
	log.Println("Green on. (real call)")
	g.green.High()
}

func (g *GpioController) YellowOn() {
	log.Println("Yellow on. (real call)")
	g.yellow.High()
}

func (g *GpioController) RedOn() {
	log.Println("Red on. (real call)")
	g.red.High()
}

func (g *GpioController) YellowRedOn() {
	log.Println("Yellow-Red on. (real call)")
	g.YellowOn()
	g.RedOn()
}

func (g *GpioController) AllOn() {
	log.Println("All on. (real call)")
	g.GreenOn()
	g.YellowRedOn()
}

func (g *GpioController) GreenOff() {
	log.Println("Green off. (real call)")
	g.green.Low()
}

func (g *GpioController) YellowOff() {
	log.Println("Yellow off. (real call)")
	g.yellow.Low()
}

func (g *GpioController) RedOff() {
	log.Println("Red off. (real call)")
	g.red.Low()
}

func (g *GpioController) YellowRedOff() {
	log.Println("Yellow-Red off. (real call)")
	g.YellowOff()
	g.RedOff()
}

func (g *GpioController) AllOff() {
	log.Println("All off. (real call)")
	g.YellowRedOff()
	g.RedOff()
}
