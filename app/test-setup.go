package app

import (
	"drones-dashboard/models"

	uuid "github.com/satori/go.uuid"
)

//droneMap will serve as our in-memory storage
var mockStore = make(map[uuid.UUID]*models.Drone)

var drone1 *models.Drone
var drone2 *models.Drone

func init() {

	id1 := uuid.NewV4()
	lat1 := 83.1424
	lon1 := -52.8805
	speed1 := 0.0
	stationary1 := true

	id2 := uuid.NewV4()
	lat2 := 42.0967
	lon2 := -141.2985
	speed2 := 0.0
	stationary2 := true

	drone1 = &models.Drone{ID: id1, Lat: lat1, Lon: lon1, Speed: speed1, Stationary: stationary1}
	drone2 = &models.Drone{ID: id2, Lat: lat2, Lon: lon2, Speed: speed2, Stationary: stationary2}

}

//Mock DB struct
type mockDB struct {
	//embedding models.Datastore makes mockDB satsify it. this way don't need to stub each method on it
	models.Datastore
}

//GetDrones sends basic drone data
func (*mockDB) GetDrones() []*models.Drone {
	drones := []*models.Drone{}
	drones = append(drones, drone1)
	drones = append(drones, drone2)

	return drones
}

//AddDrone adds drone to mockStore
func (*mockDB) AddDrone(drone *models.Drone) error {
	mockStore[drone.ID] = drone
	return nil
}
