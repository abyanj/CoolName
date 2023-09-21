package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Info struct {
	Message string `json:"message"`
}

func main() {
	// Existing route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	// New API endpoint
	http.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		info := Info{Message: "This is some information from the Golang server."}

		json.NewEncoder(w).Encode(info)
	})

	http.ListenAndServe(":8080", nil)
}
