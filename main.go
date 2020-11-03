package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	Route "./BusController"
	"github.com/gorilla/mux"
)

//var buses []Datamodel.NewBus

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", Route.HomePage).Methods("GET")
	r.HandleFunc("/details", Route.GetBus).Methods("GET")
	r.HandleFunc("/add", Route.AddBus).Methods("POST")
	r.HandleFunc("/update", Route.UpdateBus).Methods("PUT")
	r.HandleFunc("/delete", Route.DeleteBus).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}
func main() {
	file, err := os.OpenFile("Bus.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	fmt.Println("server running at localhost:5000")
	handleRequests()
}
