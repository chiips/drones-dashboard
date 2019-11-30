package app

import (
	"net/http"
)

//Server struct to include our dependencies
type Server struct {
	Router *http.ServeMux
	DB     models.Datastore
	Log    *logs.Log
}
