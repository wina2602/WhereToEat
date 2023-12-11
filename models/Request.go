package models

type Request struct {
	Id        int     `json:"id"`
	Username  string  `json:"uname"`
	Lat       float32 `json:"lat"`
	Long      float32 `json:"Long"`
	Radius    int     `json:"radius"`
	PlaceType string  `json:"type"`
}
type Saverequest struct {
	Username string `json:"uname"`
	Password string `json:"pwd"`
}
type Swiperequest struct {
	SID string `json:"sid"`
	RID string `json:"rid"`
}
type JoinroomRequest struct {
	SID      string `json:"sid"`
	Username string `json:"userName"`
}
