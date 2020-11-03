package Datamodel

type NewBus struct {
	Bus_Id      int    `json:"bus_id`
	Bus_Name    string `json:"bus_name"`
	Bus_Type    string `json:"bus_type"`
	Originate   string `json:"originate"`
	Destination string `json:"destination"`
}
