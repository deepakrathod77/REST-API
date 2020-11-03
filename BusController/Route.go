package Route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./Datamodel"
	cs "./Datamodel/cassandra/utils"

	"github.com/gocql/gocql"
)

var buses []Datamodel.NewBus

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Paytm Bus!\n")
	log.Println("visited the home page of Paytm Bus")
}

var Bus Datamodel.NewBus

//ADD NEW BUS DETAILS
func AddBus(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &Bus)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode("An error occurred. Please check  request")
		return
	}
	fmt.Println(" Adding new bus details.....\n", Bus)
	if err := cs.GetSession().Query("INSERT INTO bus(bus_id, bus_name,bus_type,originate,destination) VALUES(?, ?, ?, ?,?)",
		Bus.Bus_Id, Bus.Bus_Name, Bus.Bus_Type, Bus.Originate, Bus.Destination).Consistency(gocql.One).Exec(); err != nil {
		log.Println("Error while inserting Bus")
		log.Println(err)
	}
	log.Println("New bus details has been added..bus id:", Bus.Bus_Id)
}

//DELETE BUS DETAILS
func DeleteBus(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &Bus)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode("An error occurred. Please check  request")
		return
	}

	if err := cs.GetSession().Query("DELETE FROM bus WHERE bus_id = ?", Bus.Bus_Id).Consistency(gocql.One).Exec(); err != nil {
		log.Println("Error while deleting bus")
		log.Println(err)
	}
	log.Println(" Deleting bus details.....")
	log.Println("Bus details has been deleted....bus id:", Bus.Bus_Id)
}

//UPDATE BUS DETAILS
func UpdateBus(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	body, err := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &Bus)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		json.NewEncoder(w).Encode("An error occurred. Please check  request")
		return
	}
	if err := cs.GetSession().Query("UPDATE bus SET bus_name = ?,bus_type=? ,originate = ?,destination = ? WHERE bus_id = ?",
		Bus.Bus_Name, Bus.Bus_Type, Bus.Originate, Bus.Destination, Bus.Bus_Id).Consistency(gocql.One).Exec(); err != nil {
		log.Println("Error while updating bus")
		log.Println(err)
	}
	log.Println(" Updating bus details.....")
	log.Println("Bus details has been updated....bus id:", Bus.Bus_Id)
}

//DISPLAY BUS DETAILS
func GetBus(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all Bus information")
	m := map[string]interface{}{}

	iter := cs.GetSession().Query("SELECT * FROM bus").Consistency(gocql.One).Iter()
	for iter.MapScan(m) {
		buses = append(buses, Datamodel.NewBus{
			Bus_Id:      m["bus_id"].(int),
			Bus_Name:    m["bus_name"].(string),
			Bus_Type:    m["bus_type"].(string),
			Originate:   m["originate"].(string),
			Destination: m["destination"].(string),
		})
		m = map[string]interface{}{}
	}
	json.NewEncoder(w).Encode(buses)
}
