package main

import (
	"encoding/json"
	"net/http"
	"time"
)

// Stock ...
type Stock struct {
	Sticker string
}

// Price ...
type Price struct {
	Price     string
	Timestamp time.Time
}

func health(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
}

func stockprice(response http.ResponseWriter, request *http.Request) {

	stockpriceMap := make(map[string]string)

	stockpriceMap["COF"] = "90$"
	stockpriceMap["CNB"] = "75$"

	var s Stock
	var p Price

	err := json.NewDecoder(request.Body).Decode(&s)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	price := stockpriceMap[s.Sticker]

	response.Header().Set("Content-Type", "application/json")
	if len(price) == 0 {
		json.NewEncoder(response).Encode(p)
		return
	}
	p.Price = price
	p.Timestamp = time.Now()
	json.NewEncoder(response).Encode(p)
}

func main() {
	http.HandleFunc("/", health)
	http.HandleFunc("/stockpriceapi", stockprice)
	http.ListenAndServe(":8092", nil)
}
