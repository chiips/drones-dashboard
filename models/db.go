package models

import (
	"database/sql"

	//pq is necessary for connecting with PostgreSQL
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

//Datastore is our interface to work witih the DB
type Datastore interface {

	//Drone methods
	GetDrones() ([]*Drone, error)

}

//DB is our database type
type DB struct {
	*sql.DB
}

//NewDB creates a new DB instance
func NewDB(dbURL string) (*DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
