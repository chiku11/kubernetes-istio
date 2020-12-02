package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

// Weather ...
type Weather struct {
	Celsius    string
	Fahrenheit string
	Version    string
}

// Location ...
type Location struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func health(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
}

func findLocation(response http.ResponseWriter, request *http.Request) {

	weatherMap := make(map[string]string)

	weatherMap["elgin IL"] = "39.2f 4c"
	weatherMap["oakland CA"] = "82.4f 28c"

	var l Location
	var w Weather

	err := json.NewDecoder(request.Body).Decode(&l)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	split := strings.Split(weatherMap[l.City+" "+l.State], " ")

	response.Header().Set("Content-Type", "application/json")
	if len(split) <= 1 {
		json.NewEncoder(response).Encode(w)
		return
	}
	w.Fahrenheit = split[0]
	w.Celsius = split[1]
	w.Version = "v2"
	json.NewEncoder(response).Encode(w)
}

func main() {
	http.HandleFunc("/", health)
	http.HandleFunc("/weatherapi", findLocation)
	http.ListenAndServe(":8091", nil)
}
