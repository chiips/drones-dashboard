package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"drones-dashboard/models"
	"drones-dashboard/websock"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

//drones handler attaches to server struct so we can access all dependencies
func (s *Server) drones() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//only allow http get requests
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}

		// Upgrade the connection from a standard HTTP connection to a websocket one
		ws, err := websock.Upgrade(w, r)
		if err != nil {
			s.Log.Errorln(err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		defer ws.Close()

		//get the context of the request
		//if request cancelled by user, we can cancel our operations
		ctx := r.Context()

		//make a channel to send and receive our drones result
		droneCh := make(chan []*models.Drone)
		//make a channel to send and receive errors
		errCh := make(chan error)

		//set update time
		updateTime := 5

		//go run this anonymous function in an independent thread
		go func() {

			//send first
			drones := s.DB.GetDrones()
			if err != nil {
				errCh <- err
				return
			}
			droneCh <- drones

			// start for loop that runs for duration of our websockets connection
			for {

				//check if context cancelled before time to get drones
				//return if so
				if ctx.Err() != nil {
					return
				}

				// create a new ticker that ticks every 5 seconds
				ticker := time.NewTicker(time.Duration(updateTime) * time.Second)

				// every time our ticker ticks
				for range ticker.C {

					if ctx.Err() != nil {
						return
					}

					//call get drones from our client
					drones := s.DB.GetDrones()

					//if error returned, then send error on error channel and return from for loop
					if err != nil {
						errCh <- err
						return
					}

					//by default send the drones ont the drone channel
					droneCh <- drones

				}
			}

		}()

		//for as long as this connection stays open
		for {

			select {
			//if the context is done (because error) then return
			case <-ctx.Done():
				s.Log.Errorln(ctx.Err())
				http.Error(w, "We could not process your request at this time. Please try again later.", http.StatusRequestTimeout)
				return
			//if error sent over error channel then return
			case err := <-errCh:
				s.Log.Errorln(err)
				http.Error(w, http.StatusText(500), http.StatusInternalServerError)
				return
			//if drones sent over drone channel then send as json data
			case drones := <-droneCh:
				// marshal response into a json string
				response, err := json.Marshal(drones)
				if err != nil {
					s.Log.Errorln(err)
					http.Error(w, http.StatusText(500), http.StatusInternalServerError)
					return
				}
				// write json to websocket connection
				if err := ws.WriteMessage(websocket.TextMessage, []byte(response)); err != nil {
					if err != nil {
						s.Log.Errorln(err)
						http.Error(w, http.StatusText(500), http.StatusInternalServerError)
						return
					}
				}
			}

		}

	}
}

func (s *Server) add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
		}

		submission := &models.Drone{}
		err := json.NewDecoder(r.Body).Decode(&submission)
		if err != nil {
			s.Log.Errorln(err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		//create uuid for post
		submission.ID = uuid.NewV4()
		submission.Speed = 0
		submission.Stationary = true

		ctx := r.Context()
		okCh := make(chan bool)
		errCh := make(chan error)

		go func() {

			//check if context cancelled before time to talk to DB
			if ctx.Err() != nil {
				return
			}

			err := s.DB.AddDrone(submission)

			if err != nil {
				errCh <- err
				return
			}

			okCh <- true
			return

		}()

		select {
		case <-ctx.Done():
			s.Log.Errorln(ctx.Err())
			http.Error(w, "We could not process your request at this time. Please try again later.", http.StatusRequestTimeout)
			return
		case err := <-errCh:
			s.Log.Errorln("error submitting post:", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		case <-okCh:
			fmt.Fprint(w, "drone added!\n")
			return
		}

	}

}

func (s *Server) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodDelete {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}

		query, ok := r.URL.Query()["id"]

		if !ok || len(query[0]) <= 1 || strings.TrimSpace(query[0]) == "" {
			s.Log.Errorln("no id param")
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}

		id, err := uuid.FromString(query[0])
		if err != nil {
			s.Log.Errorln(err)
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
			return
		}

		okCh := make(chan bool)
		errCh := make(chan error)

		ctx := r.Context()

		go func() {

			//check if context cancelled before time to talk to DB
			if ctx.Err() != nil {
				return
			}

			err = s.DB.DeleteDrone(id)

			if err != nil {
				errCh <- err
				return
			}

			okCh <- true
			return

		}()

		select {
		case <-ctx.Done():
			s.Log.Errorln(ctx.Err())
			http.Error(w, "We could not process your request at this time. Please try again later.", http.StatusRequestTimeout)
			return
		case err := <-errCh:
			s.Log.Errorln("error deleting drone:", err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		case <-okCh:
			fmt.Fprint(w, "drone deleted!")
			return
		}

	}
}
