package app

import (
	"bytes"
	"drones-dashboard/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
)

func TestAdd(t *testing.T) {

	//initialize router and server
	router := http.NewServeMux()
	s := Server{DB: &mockDB{}, Router: router}
	s.Routes()

	//set up json
	newDrone := []byte(`{"lat":43,"lon":79}`)

	//set up recorder and request including json
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(newDrone))
	if err != nil {
		t.Fatal(err)
	}

	//serve the request
	router.ServeHTTP(rr, req)

	//check the status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code:\ngot:\n%v\n want:\n%v", status, http.StatusOK)
	}
}

//TestDrones tests the drones handler
func TestDrones(t *testing.T) {

	//create a new router
	router := http.NewServeMux()
	//create a server struct with our mockclient and router
	s := Server{DB: &mockDB{}, Router: router}
	//init all routes
	s.Routes()

	//create a test server
	srv := httptest.NewServer(s.Router)
	defer srv.Close()

	//define websocket url
	u := "ws" + strings.TrimPrefix(srv.URL, "http") + "/drones"

	//dial into the websocket
	ws, resp, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	//ensure we receive the appropriate statuscode
	if got, want := resp.StatusCode, http.StatusSwitchingProtocols; got != want {
		t.Errorf("resp.StatusCode = %q, want %q", got, want)
	}

	//ensure we can read messages from the websocket
	_, p, err := ws.ReadMessage()
	if err != nil {
		t.Fatalf("%v", err)
	}

	//read json data from the websocket into our models.Drone object for comparing
	got := []*models.Drone{}
	err = json.Unmarshal(p, &got)
	if err != nil {
		t.Fatalf("%v", err)
	}

	//set our expected drones
	want := []*models.Drone{}
	want = append(want, drone1)
	want = append(want, drone2)

	//ensure what we got matches what we want
	same, gw := compareDroneList(got, want)

	if !same {
		t.Error(gw)
	}

}

//compareDroneList ranges over our result and compares the values of each drone to the values of our expected drones
func compareDroneList(got, want []*models.Drone) (bool, string) {
	for key, droneGot := range got {
		if droneGot.ID != want[key].ID {
			gw := fmt.Sprintf("handler returned wrong drone ID:\ngot: %v\nwant: %v", droneGot.ID, want[key].ID)
			return false, gw
		}
		if droneGot.Lat != want[key].Lat {
			gw := fmt.Sprintf("handler returned wrong drone ID:\ngot: %v\nwant: %v", droneGot.Lat, want[key].Lat)
			return false, gw
		}
		if droneGot.Lon != want[key].Lon {
			gw := fmt.Sprintf("handler returned wrong drone ID:\ngot: %v\nwant: %v", droneGot.Lon, want[key].Lon)
			return false, gw
		}
		if droneGot.Speed != want[key].Speed {
			gw := fmt.Sprintf("handler returned wrong drone ID:\ngot: %v\nwant: %v", droneGot.Speed, want[key].Speed)
			return false, gw
		}
		if droneGot.Stationary != want[key].Stationary {
			gw := fmt.Sprintf("handler returned wrong drone ID:\ngot: %v\nwant: %v", droneGot.Stationary, want[key].Stationary)
			return false, gw
		}
	}
	return true, ""
}
