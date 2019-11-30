package models


//Coords type defined
type Coords struct {
	ID  uuid.UUID `json:"id"`
	Lat float64   `json:"lat"`
	Lon float64   `json:"lon"`
}