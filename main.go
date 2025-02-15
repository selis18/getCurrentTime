package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time     string `json:"time"`
	Timezone string `json:"timezone"`
}

func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8080", nil)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	timezone := r.URL.Query().Get("timezone")

	if timezone == "" {
		timezone = "UTC"
	}

	location, err := time.LoadLocation(timezone)
	if err != nil {
		http.Error(w, "Invalid timezone", http.StatusBadRequest)
		return
	}

	currentTime = currentTime.In(location)

	response := TimeResponse{
		Time:     currentTime.Format("02-01-2006 15:04:05"),
		Timezone: timezone,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
