package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type weatherData struct {
	LocationName string   `json: locationName`
	Weather      string   `json: weather`
	Temperature  int      `json: temperature`
	Celsius      bool     `json: celsius`
	TempForecast []int    `json: temp_forecast`
	Wind         windData `json: wind`
}

type windData struct {
	Direction string `json: direction`
	Speed     int    `json: speed`
}

type loc struct {
	Lat float32 `json: lat`
	Lon float32 `json: lon`
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	location := loc{}

	log.Println(r.Method)
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error reading the body", err)
	}

	err = json.Unmarshal(jsn, &location)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	log.Printf("Received: %v\n", location)
	weather := weatherData{
		LocationName: "Zzyzx",
		Weather:      "cloudy",
		Temperature:  31,
		Celsius:      true,
		TempForecast: []int{30, 32, 29},
		Wind: windData{
			Direction: "S",
			Speed:     20,
		},
	}

	weatherJson, err := json.Marshal(weather)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(weatherJson)

}

func server() {
	http.HandleFunc("/", weatherHandler)
	http.ListenAndServe(":8080", nil)
}

func client() {
	locJson, err := json.Marshal(loc{Lat: 35.14326, Lon: -116.104})
	req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer(locJson))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(body))
	resp.Body.Close()
}

func main() {
	go server()
	client()
}
