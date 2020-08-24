package lights

// TrafficLight can be operated like a traffic light.
type TrafficLight interface {
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

// NewTrafficLight creates a new traffic light with the passed type an pin configuration.
func NewTrafficLight(controllerType string, greenPin int, yellowPin int, redPin int) TrafficLight {
	switch controllerType {
	case "demo":
		return DemoTrafficLight{}
	case "negated":
		return NewNegatedTrafficLight(greenPin, yellowPin, redPin)
	case "classic":
		return NewClassicTrafficLight(greenPin, yellowPin, redPin)
	default:
		return DemoTrafficLight{}
	}
}
