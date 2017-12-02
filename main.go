package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorldEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloWorldEndpoint).Methods("GET")
	/*router.HandleFunc("/join", HelloWorldEndpoint).Methods("GET")
	router.HandleFunc("/game/{id}/slot/{slotIndex}", HelloWorldEndpoint).Methods("PUT")
	router.HandleFunc("/game/{id}/status", HelloWorldEndpoint).Methods("GET")
	router.HandleFunc("/game/{id}", HelloWorldEndpoint).Methods("GET")*/

	log.Fatal(http.ListenAndServe(":8080", router))
}
