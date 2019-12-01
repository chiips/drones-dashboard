package models

import (
	"math"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

//Approximate coordinates of Canadian border limits
var maxLat = 83.1424
var minLat = 42.0967
var maxLon = -52.8805
var minLon = -141.2985

const updateTime = 10

func init() {

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	//create five drones
	for i := 1; i <= 5; i++ {

		id := uuid.NewV4()

		lat := minLat + r.Float64()*(maxLat-minLat)
		lon := minLon + r.Float64()*(maxLon-minLon)

		drone := &Drone{
			ID:  id,
			Lat: lat,
			Lon: lon,
		}

		droneStore[drone.ID] = drone
	}

	//go run update function in an indepedent thread
	go update()

}

//update function mimics drones sending new data to the server
//every 10 seconds, as long as the program runs, the function will update the position of each drone
func update() {
	for {
		// we create a new ticker that ticks every 10 seconds
		ticker := time.NewTicker(10 * time.Second)

		s := rand.NewSource(time.Now().UnixNano())
		r := rand.New(s)

		// every time our ticker ticks
		for range ticker.C {
			//move each drone
			for _, drone := range droneStore {

				latChange := 0 + r.Float64()*(0.0010-0)
				lonChange := 0 + r.Float64()*(0.0010-0)

				change := r.Intn(1)

				oldLat := drone.Lat
				oldLon := drone.Lon

				if change == 1 {
					drone.Lat += latChange
					drone.Lon += lonChange
				} else {
					drone.Lat -= latChange
					drone.Lon -= lonChange
				}

				//calculate distance using Haversine formula
				dist := distance(oldLat, drone.Lat, oldLon, drone.Lon)

				//distance in meters. if moved less than one meter in 10 seconds (since data updated and sent every 10 seconds) then count as stationary
				if dist <= 1 {
					drone.Stationary = true
				} else {
					drone.Stationary = false
				}

				speed := dist / updateTime
				drone.Speed = math.Round(speed*100) / 100

			}

		}
	}
}

//distance uses Harversine formula to calculate distance according to latitude and longitude of drones
func distance(lat1, lat2, lon1, lon2 float64) float64 {

	// convert to radians
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	// Earth radius in meters
	r = 6378100

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))

}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}