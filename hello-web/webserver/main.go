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
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	http.ListenAndServe(":8000", r)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	name := v.Get("name")
	url := fmt.Sprintf("http://hello.api/hello?name=%s", name)
	response, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&APIError{
			Message: err.Error(),
		})
	}
	w.Header().Set("Content-Type", "application/json")
	data, _ := ioutil.ReadAll(response.Body)
	g := &Greeting{}
	err = json.Unmarshal(data, &g)
	json.NewEncoder(w).Encode(g)
}

// Greeting response model
type Greeting struct {
	Message string `json:"greeting"`
}

func getWeather(w http.ResponseWriter, r *http.Request) {
	zip := mux.Vars(r)["zip"]
	url := fmt.Sprintf("http://hello.api/weather/%s", zip)
	response, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&APIError{
			Message: err.Error(),
		})
	}
	data, _ := ioutil.ReadAll(response.Body)
	forecast := &Forecast{}
	err = json.Unmarshal(data, &forecast)
	json.NewEncoder(w).Encode(forecast)

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
