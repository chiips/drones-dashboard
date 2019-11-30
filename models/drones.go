package models


//Drone type defined
type Drone struct {
	ID         uuid.UUID `json:"id"`
	Lat        float64   `json:"lat"`
	Lon        float64   `json:"lon"`
	Speed      float64   `json:"speed"`
	Stationary bool      `json:"stationary"`
}

