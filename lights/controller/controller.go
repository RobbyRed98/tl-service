package controller

type Controller interface {
	GreenOn()
	YellowOn()
	RedOn()
	YellowRedOn()
	AllOn()
	GreenOff()
	YellowOff()
	RedOff()
	YellowRedOff()
	AllOff()
}

func NewController(controllerType string, greenPin int, yellowPin int, redPin int) Controller {
	switch controllerType {
	case "demo":
		return DemoController{}
	case "negated":
		return NewNegatedGpioController(greenPin, yellowPin, redPin)
	case "classic":
		return NewGpioController(greenPin, yellowPin, redPin)
	default:
		return DemoController{}
	}
}
