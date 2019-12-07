package models

import uuid "github.com/satori/go.uuid"

//Datastore is our interface to work witih the DB
type Datastore interface {
	GetDrones() []*Drone
	AddDrone(drone *Drone) error
	DeleteDrone(id uuid.UUID) error
}

//DB is our database type
type DB struct{}

//NewDB creates a new DB instance
func NewDB() *DB { return &DB{} }
