package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/RobbyRed98/tl-service/config"
	"github.com/RobbyRed98/tl-service/handler"
	"github.com/RobbyRed98/tl-service/lights"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	host := flag.String("ip", "localhost", "The host address to run the tl-service on.")
	port := flag.Int("port", 8080, "The port to run the tl-service on.")
	green := flag.Int("green", 20, "GPIO the green light is connected to.")
	yellow := flag.Int("yellow", 21, "GPIO the yellow light is connected to.")
	red := flag.Int("red", 20, "GPIO the red light is connected to.")
	typeValue := flag.String("type", "demo", "The controller type to use. (demo|classic|negated)")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGQUIT)
	go func() {
		select {
		case <-sigs:
			log.Println("Exiting application.")
			os.Exit(0)
		}
	}()

	flag.Parse()

	log.Println("Starting traffic light controller service...")

	var conf config.TrafficLightConfig
	aController := lights.NewTrafficLight(*typeValue, *green, *yellow, *red)
	handler := handler.NewTrafficLightHandler(conf, aController)
	router := mux.NewRouter()

	lights.ShowLights(aController)

	router.HandleFunc("/", handler.GetAvailablity).Methods("GET")
	router.HandleFunc("/heartbeat", handler.GetHeartbeat).Methods("GET")
	router.HandleFunc("/config", handler.GetConfig).Methods("GET")
	router.HandleFunc("/start", handler.StartTrafficLights).Methods("GET")
	router.HandleFunc("/stop", handler.StopTrafficLights).Methods("GET")
	router.HandleFunc("/running", handler.GetCurrentState).Methods("GET")
	router.HandleFunc("/config", handler.CreateConfig).Methods("POST")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), handlers.CORS()(router)))
}
