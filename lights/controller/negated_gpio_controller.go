package controller

import (
	"log"
)

type NegatedGpioController struct {
	gpioController GpioController
}

func NewNegatedGpioController(greenPinId int8, yellowPinId int8, redPinId int8) *NegatedGpioController {
	n := new(NegatedGpioController)
	n.gpioController = *NewGpioController(greenPinId, yellowPinId, redPinId)
	return n
}

func (n *NegatedGpioController) GreenOn() {
	log.Println("Green on. (negated call)")
	n.gpioController.GreenOff()
}

func (n *NegatedGpioController) YellowOn() {
	log.Println("Yellow on. (negated call)")
	n.gpioController.YellowOff()
}

func (n *NegatedGpioController) RedOn() {
	log.Println("Red on. (negated call)")
	n.gpioController.RedOff()
}

func (n *NegatedGpioController) YellowRedOn() {
	log.Println("Yellow-Red on. (negated call)")
	n.gpioController.YellowRedOff()
}

func (n *NegatedGpioController) AllOn() {
	log.Println("All on. (negated call)")
	n.gpioController.AllOff()
}

func (n *NegatedGpioController) GreenOff() {
	log.Println("Green off. (negated call)")
	n.gpioController.GreenOn()
}

func (n *NegatedGpioController) YellowOff() {
	log.Println("Yellow off. (negated call)")
	n.gpioController.YellowOn()
}

func (n *NegatedGpioController) RedOff() {
	log.Println("Red off. (negated call)")
	n.gpioController.RedOn()
}

func (n *NegatedGpioController) YellowRedOff() {
	log.Println("Yellow-Red off. (negated call)")
	n.gpioController.YellowRedOn()
}

func (n *NegatedGpioController) AllOff() {
	log.Println("All off. (negated call)")
	n.gpioController.AllOn()
}
