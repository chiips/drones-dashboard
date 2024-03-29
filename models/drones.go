package models

import (
	uuid "github.com/satori/go.uuid"
)

//Drone type defined
type Drone struct {
	ID         uuid.UUID `json:"id"`
	Lat        float64   `json:"lat"`
	Lon        float64   `json:"lon"`
	Speed      float64   `json:"speed"`
	Stationary bool      `json:"stationary"`
}

//droneMap will serve as our in-memory storage
var droneStore = make(map[uuid.UUID]*Drone)

//GetDrones sends basic drone data
func (*DB) GetDrones() []*Drone {
	result := []*Drone{}

	for _, drone := range droneStore {
		result = append(result, drone)
	}

	return result
}

//AddDrone adds drone to DroneStore
func (*DB) AddDrone(drone *Drone) error {
	droneStore[drone.ID] = drone
	return nil
}

//DeleteDrone removes drone from map
func (*DB) DeleteDrone(id uuid.UUID) error {
	delete(droneStore, id)
	return nil
}
