package lights

const (
	// Demo type of the traffic light.
	Demo string = "demo"
	// Classic type of the traffic light.
	Classic string = "classic"
	// Negated type of the traffic light.
	Negated string = "negated"
)



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
func NewTrafficLight(targetType string, greenPin int, yellowPin int, redPin int) TrafficLight {
	switch targetType {
	case Demo:
		return DemoTrafficLight{}
	case Negated:
		return NewNegatedTrafficLight(greenPin, yellowPin, redPin)
	case Classic:
		return NewClassicTrafficLight(greenPin, yellowPin, redPin)
	default:
		return DemoTrafficLight{}
	}
}
