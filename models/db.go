package models

//Datastore is our interface to work witih the DB
type Datastore interface {
	GetDrones() []*Drone
	AddDrone(drone *Drone) error
}

//DB is our database type
type DB struct{}

//NewDB creates a new DB instance
func NewDB() *DB { return &DB{} }
