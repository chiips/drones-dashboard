package app

import (
	"net/http"

	"dashboard/backend/models"
	"dashboard/backend/logs"
)

//Server struct to include our dependencies
type Server struct {
	Router *http.ServeMux
	DB     models.Datastore
	Log    *logs.Log
}
