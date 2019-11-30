package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

func (s *Server) add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(400), http.StatusBadRequest)
		}

		submission := &models.DroneData{}
		err := json.NewDecoder(r.Body).Decode(&submission)
		if err != nil {
			s.Log.Errorln(err)
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		//create uuid for post
		id := uuid.NewV4()
		// if err != nil {
		// 	s.Log.Errorln(err)
		// 	http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		// 	return
		// }

		submission.ID = id

		central.DroneStore = append(central.DroneStore, submission)

		fmt.Println("added")
	}
}
