package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RobbyRed98/tl-service/config"
	"github.com/RobbyRed98/tl-service/lights"
)

// TrafficLightHandler handles the incoming requests and operates the traffic light according to these.
type TrafficLightHandler struct {
	config           *config.TrafficLightConfig
	channel          chan string
	lightsRunning    bool
	lightsController lights.TrafficLight
}

// NewTrafficLightHandler returns a new TrafficLightHandler
func NewTrafficLightHandler(conf config.TrafficLightConfig, lightsController lights.TrafficLight) *TrafficLightHandler {
	t := new(TrafficLightHandler)
	t.config = &conf
	t.channel = make(chan string)
	t.lightsRunning = false
	t.lightsController = lightsController
	return t
}

// Availablity shows the visual in the browser availablity by responding.
func (t *TrafficLightHandler) Availablity(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Traffic Light Controller is available. (GO)</p>")
}

// Config returns the current configuration as JSON.
func (t *TrafficLightHandler) Config(w http.ResponseWriter, r *http.Request) {
	log.Println(t.config)

	if !t.config.IsValid() {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(t.config)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// UpdateConfig updates the current configuration used to operate the TrafficLightHandler.
func (t *TrafficLightHandler) UpdateConfig(w http.ResponseWriter, r *http.Request) {
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

// Heartbeat provides a heartbeat signal which ensures the system is running.
func (t *TrafficLightHandler) Heartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// StartTrafficLight triggers the start of the traffic light.
func (t *TrafficLightHandler) StartTrafficLight(w http.ResponseWriter, r *http.Request) {
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

// StopTrafficLight triggers the stop of the traffic light.
func (t *TrafficLightHandler) StopTrafficLight(w http.ResponseWriter, r *http.Request) {
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

// CurrentState returns if the traffic light is running.
func (t *TrafficLightHandler) CurrentState(w http.ResponseWriter, r *http.Request) {
	if t.lightsRunning {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
