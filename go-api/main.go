package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	state := func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}

	http.HandleFunc("/tesla/engine", state)

	fmt.Println(`curl -H "Content-Type: application/Json" -XPOST -d '{"status": "STARTED", "id": "a12163bf-6e33-57d8-959c-559f5ca44a22"}' 127.0.0.1:8080/tesla/engine`)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
