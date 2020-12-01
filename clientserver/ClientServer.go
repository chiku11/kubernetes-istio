package main

import (
	"io/ioutil"
	"net/http"
)

func health(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
}

func weather(w http.ResponseWriter, req *http.Request) {
	response, err := http.Post("http://host.docker.internal:80/weatherapi", "application/json", req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func stockprice(w http.ResponseWriter, req *http.Request) {
	response, err := http.Post("http://host.docker.internal:80/stockpriceapi", "application/json", req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func main() {
	http.HandleFunc("/", health)
	http.HandleFunc("/weather", weather)
	http.HandleFunc("/stock", stockprice)
	http.ListenAndServe(":8090", nil)
}
