package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"dashboard/backend/app"
	"dashboard/backend/logs"
	"dashboard/backend/models"
)

func main() {

	//set up new logger (currently no .env file specifying logfile, so defaults to terminal output)
	logfile := os.Getenv("logfile")
	logger, err := logs.NewLogger(logfile)
	if err != nil {
		fmt.Println("error setting up new logger:", err)
	}

	//set up new router
	router := http.NewServeMux()

	db := models.NewDB()

	//assign router, drone client, logger to our app's Server struct
	s := app.Server{Router: router, DB: db, Log: logger}
	//init Server routes
	s.Routes()

	//define port
	port := ":8080"

	//define http server
	srv := &http.Server{
		Addr:         port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  20 * time.Second,
		Handler:      s.Router,
	}

	//listen
	s.Log.Infoln("Listening on:", port)
	s.Log.Fatalln(srv.ListenAndServe())

}
