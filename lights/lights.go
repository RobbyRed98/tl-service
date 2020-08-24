package lights

import (
	"fmt"
	"log"
	"time"

	"github.com/RobbyRed98/tl-service/config"
)

const (
	green int = iota
	yellow
	red
	yellowRed
	startRed
)

// ShowLights turns all traffic lights on for three seconds to show they are functional.
func ShowLights(lightsController TrafficLight) {
	lightsController.AllOn()
	stopChannel := time.Tick(3 * time.Second)
	select {
	case <-stopChannel:
		lightsController.AllOff()
		return
	}
}

// Stop turns the traffic light off.
func Stop(channel chan string) {
	channel <- "stop"
}

// Start turns the traffic light on.
func Start(channel chan string, conf config.TrafficLightConfig, lightsController TrafficLight) {
	log.Println(conf)
	nextState := startRed
	timer := *time.NewTimer(time.Duration(1000000000 * conf.RandRed()))
	for {
		select {
		case <-timer.C:
			switch nextState {
			case startRed:
				lightsController.RedOn()
				nextState = yellowRed
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.RandRed()))
				timer.Reset(timeout)
				break
			case red:
				lightsController.YellowOff()
				lightsController.RedOn()
				nextState = yellowRed
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.RandRed()))
				timer.Reset(timeout)
				break
			case yellowRed:
				lightsController.RedOff()
				lightsController.YellowRedOn()
				nextState = green
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.YellowRed))
				timer.Reset(timeout)
				break
			case green:
				lightsController.YellowRedOff()
				lightsController.GreenOn()
				nextState = yellow
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.Green))
				timer.Reset(timeout)
				break
			case yellow:
				lightsController.GreenOff()
				lightsController.YellowOn()
				nextState = red
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.Yellow))
				timer.Reset(timeout)
				break
			}
		case <-channel:
			log.Println("Stopping previous light controlling routine.")
			lightsController.AllOff()
			return
		}
	}
}
