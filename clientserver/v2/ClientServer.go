package main

import (
	"io/ioutil"
	"net/http"
)

func health(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
}

func weather(w http.ResponseWriter, req *http.Request) {
	invokeAPI(w, req, "http://host.docker.internal:80/weatherapi")
}

func stockprice(w http.ResponseWriter, req *http.Request) {
	invokeAPI(w, req, "http://host.docker.internal:80/stockpriceapi")
}

func invokeAPI(w http.ResponseWriter, req *http.Request, url string) {

	response, err := http.Post(url, "application/json", req.Body)
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
