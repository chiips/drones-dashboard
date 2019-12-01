package app

import (
	"net/http"

	"dashboard/backend/logs"
	"dashboard/backend/models"
)

//Server struct to include our dependencies
type Server struct {
	Router *http.ServeMux
	DB     models.Datastore
	Log    *logs.Log
}
