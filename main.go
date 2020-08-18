package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type TLConfig struct {
	Green     int32 `json:"greenLightDuration"`
	Yellow    int32 `json:"yellowLightDuration"`
	YellowRed int32 `json:"yellowRedLightDuration"`
	RedLower  int32 `json:"lowerIntervalBorder"`
	RedUpper  int32 `json:"upperIntervalBorder"`
}

func (tc TLConfig) RandRed() int32 {
	return rand.Int31n(tc.RedUpper-tc.RedLower) + tc.RedLower
}

var config TLConfig //= TLConfig{1, 1, 1, 1, 5}
var running bool = false
var stopChannel = make(chan string)

func operateTrafficLight(quit chan string, activeConfig TLConfig) {
	log.Println(activeConfig)
	state := "red"
	timer := *time.NewTimer(time.Duration(1000000000 * activeConfig.RandRed()))
	for {
		select {
		case <-timer.C:
			switch state {
			case "red":
				state = "yellowRed"
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", activeConfig.YellowRed))
				timer.Reset(timeout)
				break
			case "yellowRed":
				state = "green"
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", activeConfig.Green))
				timer.Reset(timeout)
				break
			case "green":
				state = "yellow"
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", activeConfig.Yellow))
				timer.Reset(timeout)
				break
			case "yellow":
				state = "red"
				timeout, _ := time.ParseDuration(fmt.Sprintf("%ds", activeConfig.RandRed()))
				timer.Reset(timeout)
				break
			}
			log.Println(state)
		case <-quit:
			return
		}
	}
}

func (tc TLConfig) IsValid() bool {
	if tc.Green <= 0 {
		return false
	}
	if tc.Yellow <= 0 {
		return false
	}
	if tc.YellowRed <= 0 {
		return false
	}
	if tc.RedLower <= 0 {
		return false
	}

	return tc.RedLower <= tc.RedUpper
}

func GetAvailablity(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>Traffic Light Controller is available. (GO)</p>")
}

func GetHeartbeat(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func StartTrafficLights(w http.ResponseWriter, r *http.Request) {
	if !config.IsValid() {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if running {
		stopChannel <- "stop"
	}
	go operateTrafficLight(stopChannel, config)
	running = true
	w.WriteHeader(http.StatusOK)
}

func StopTrafficLights(w http.ResponseWriter, r *http.Request) {
	if !running {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	stopChannel <- "stop"
	running = false
	w.WriteHeader(http.StatusOK)
}

func GetCurrentState(w http.ResponseWriter, r *http.Request) {
	if running {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	if !config.IsValid() {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(config)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func CreateConfig(w http.ResponseWriter, r *http.Request) {
	var newConfig TLConfig
	json.NewDecoder(r.Body).Decode(&newConfig)
	log.Println(newConfig)
	if !newConfig.IsValid() {
		log.Println("Config update failed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	config = newConfig
	log.Println("Succeessfully updated config.")
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	log.Println("Starting traffic light controller service...")

	router := mux.NewRouter()
	router.HandleFunc("/", GetAvailablity).Methods("GET")
	router.HandleFunc("/heartbeat", GetHeartbeat).Methods("GET")
	router.HandleFunc("/config", GetConfig).Methods("GET")
	router.HandleFunc("/start", StartTrafficLights).Methods("GET")
	router.HandleFunc("/stop", StopTrafficLights).Methods("GET")
	router.HandleFunc("/running", GetCurrentState).Methods("GET")
	router.HandleFunc("/config", CreateConfig).Methods("POST")
	log.Fatal(http.ListenAndServe("localhost:8080", handlers.CORS()(router)))
}
