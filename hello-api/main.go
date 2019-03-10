package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", sayHello).Methods("GET")
	r.HandleFunc("/weather/{", getWeather).Methods("GET")
	http.ListenAndServe(":5000", r)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	name := v.Get("name")
	if name == "" {
		name = "stranger"
	}
	greeting := &Greeting{
		Message: fmt.Sprintf("how ya doin, %s?", name),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(greeting)
}

// Greeting response model
type Greeting struct {
	Message string `json:"greeting"`
}

//https://weather.cit.api.here.com/weather/1.0/report.json?product=observation&zipcode=98101&oneobservation=true&app_id=DemoAppId01082013GAL&app_code=AJKnXv84fjrb0KIHawS0Tg
