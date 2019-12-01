package app

import (
	"net/http"
)

//Routes inits our Server's routes
func (s *Server) Routes() {

	//serve our index.html
	s.Router.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))

	s.Router.HandleFunc("/drones", s.drones())
	s.Router.HandleFunc("/add", s.add())

}
