package app

import (
	"net/http"

	"drones-dashboard/logs"
	"drones-dashboard/models"
)

//Server struct to include our dependencies
type Server struct {
	Router *http.ServeMux
	DB     models.Datastore
	Log    *logs.Log
}
