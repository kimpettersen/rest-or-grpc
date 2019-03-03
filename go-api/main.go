package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Tesla struct {
	Status string `json:"status"`
}

func main() {
	start := func(w http.ResponseWriter, _ *http.Request) {
		// TODO: start the car
		tesla := Tesla{
			Status: "STARTED",
		}
		data, _ := json.Marshal(tesla)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

	stop := func(w http.ResponseWriter, _ *http.Request) {
		// TODO: stop the car
		tesla := Tesla{
			Status: "STOPPED",
		}
		data, _ := json.Marshal(tesla)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

	http.HandleFunc("/tesla/123/start", start)
	http.HandleFunc("/tesla/123/stop", stop)
	fmt.Println("Start: http://localhost:8080/tesla/123/start")
	fmt.Println("Stop: http://localhost:8080/tesla/123/stop")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
