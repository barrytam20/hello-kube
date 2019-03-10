package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", sayHello).Methods("GET")
	r.HandleFunc("/weather/{zip:[0-9]{5}}", getWeather).Methods("GET")
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

func getWeather(w http.ResponseWriter, r *http.Request) {
	zip := mux.Vars(r)["zip"]
	url := fmt.Sprintf("https://weather.cit.api.here.com/weather/1.0/report.json?product=observation&zipcode=%s&oneobservation=true&app_id=DemoAppId01082013GAL&app_code=AJKnXv84fjrb0KIHawS0Tg",
		zip)
	response, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&APIError{
			Message: err.Error(),
		})
	}
	data, _ := ioutil.ReadAll(response.Body)
	weatherModel := &CITWeather{}
	err = json.Unmarshal(data, &weatherModel)
	if weatherModel.Type == "Invalid Request" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&APIError{
			Message: fmt.Sprintf("weather for %s not found", zip),
		})
	} else {
		res := &Forecast{
			Description: weatherModel.Observations.Locations[0].Forecast[0].Description,
			Icon:        weatherModel.Observations.Locations[0].Forecast[0].Icon,
			Country:     weatherModel.Observations.Locations[0].Country,
			State:       weatherModel.Observations.Locations[0].State,
			City:        weatherModel.Observations.Locations[0].City,
		}
		json.NewEncoder(w).Encode(res)
	}
}

// CITWeather model for weather response from CIT
type CITWeather struct {
	Observations CITObservations `json:"observations"`
	Type         string          `json:"Type"`
}

// CITObservations model for weather response from CIT
type CITObservations struct {
	Locations []CITLocation `json:"location"`
}

// CITLocation model for weather response from CIT
type CITLocation struct {
	Forecast []CITForecast `json:"observation"`
	Country  string        `json:"country"`
	State    string        `json:"state"`
	City     string        `json:"city"`
}

// CITForecast model for weather response from CIT
type CITForecast struct {
	Description string `json:"description"`
	Icon        string `json:"iconLink"`
}

// Forecast is the model for api response
type Forecast struct {
	Description string `json:"description"`
	Icon        string `json:"iconLink"`
	Country     string `json:"country"`
	State       string `json:"state"`
	City        string `json:"city"`
}

// APIError is a model for api errors
type APIError struct {
	Message string `json:"message"`
}
