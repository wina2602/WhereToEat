package models

type Room struct {
	SID       string
	Lat       float32
	Long      float32
	PlaceType string
	Radius    int
	Creator   string
}
