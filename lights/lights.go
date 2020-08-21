package lights

import (
	"fmt"
	"log"
	"time"

	"github.com/RobbyRed98/tl-service/config"
	"github.com/RobbyRed98/tl-service/lights/controller"
)

const (
	Green int = iota
	Yellow
	Red
	YellowRed
	StartRed
)

func ShowLights(lightsController controller.Controller) {
	lightsController.AllOn()
	stopChannel := time.Tick(3 * time.Second)
	select {
	case <-stopChannel:
		lightsController.AllOff()
		return
	}
}

func Stop(channel chan string) {
	channel <- "stop"
}

func Start(channel chan string, conf config.TrafficLightConfig, lightsController controller.Controller) {
	log.Println(conf)
	nextState := StartRed
	timer := *time.NewTimer(time.Duration(1000000000 * conf.RandRed()))
	for {
		select {
		case <-timer.C:
			switch nextState {
			case StartRed:
				lightsController.RedOn()
				nextState = YellowRed
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.RandRed()))
				timer.Reset(timeout)
				break
			case Red:
				lightsController.YellowOff()
				lightsController.RedOn()
				nextState = YellowRed
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.RandRed()))
				timer.Reset(timeout)
				break
			case YellowRed:
				lightsController.RedOff()
				lightsController.YellowRedOn()
				nextState = Green
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.YellowRed))
				timer.Reset(timeout)
				break
			case Green:
				lightsController.YellowRedOff()
				lightsController.GreenOn()
				nextState = Yellow
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", conf.Green))
				timer.Reset(timeout)
				break
			case Yellow:
				lightsController.GreenOff()
				lightsController.YellowOn()
				nextState = Red
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
