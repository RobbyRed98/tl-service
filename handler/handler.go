package handler

import(
	"encoding/json"
	"log"
	"fmt"
	"net/http"
	"github.com/RobbyRed98/tl-service/config"
	"github.com/RobbyRed98/tl-service/lights"
	"github.com/RobbyRed98/tl-service/lights/controller"
)

type TrafficLightHandler struct {
	config *config.TrafficLightConfig
	channel chan string
	lightsRunning bool
	lightsController controller.Controller
}

func NewTrafficLightHandler(conf config.TrafficLightConfig, lightsController controller.Controller) *TrafficLightHandler {
    t := new(TrafficLightHandler)
	t.config = &conf
	t.channel = make(chan string)
	t.lightsRunning = false
	t.lightsController = lightsController
    return t
}

func (t* TrafficLightHandler) GetAvailablity(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Traffic Light Controller is available. (GO)</p>")
}

func (t* TrafficLightHandler) GetConfig(w http.ResponseWriter, r *http.Request) {
	log.Println(t.config)

	if !t.config.IsValid() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(t.config)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (t* TrafficLightHandler) CreateConfig(w http.ResponseWriter, r *http.Request) {
	var newConfig config.TrafficLightConfig
	json.NewDecoder(r.Body).Decode(&newConfig)
	log.Println(newConfig)

	if !newConfig.IsValid() {
		log.Println("Config update failed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.config = &newConfig
	log.Println("Successfully updated config.")
	w.WriteHeader(http.StatusNoContent)
}

func (t* TrafficLightHandler) GetHeartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (t* TrafficLightHandler) StartTrafficLights(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting traffic lights...")
	if !t.config.IsValid() {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if t.lightsRunning {
		lights.Stop(t.channel)
		t.lightsRunning = false
	}

	go lights.Start(t.channel, *t.config, t.lightsController)
	t.lightsRunning = true

	log.Println("Started traffic lights.")

	w.WriteHeader(http.StatusOK)
}

func (t* TrafficLightHandler) StopTrafficLights(w http.ResponseWriter, r *http.Request) {
	if !t.lightsRunning {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if t.lightsRunning {
		lights.Stop(t.channel)
		t.lightsRunning = false
	}

	log.Println("Stopping traffic lights.")
	
	w.WriteHeader(http.StatusOK)
}

func (t* TrafficLightHandler) GetCurrentState(w http.ResponseWriter, r *http.Request) {
	if t.lightsRunning {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}