package messages

type Err struct {
	Message string `json:"message"`
}
type UserCreatedMsg struct {
	ID       int    `json:"id"`
	Username string `json:"uname"`
	Message  string `json:"message"`
}
type RoomCreatedMsg struct {
	Message string  `json:"message"`
	SID     string  `json:"sid"`
	Lat     float32 `json:"lat"`
	Long    float32 `json:"lang"`
	Radius  int     `json:"radius"`
	Creator string  `json:"creator"`
}
type JoinedRoomMsg struct {
	Message   string `json:"message"`
	Lat       string `json:"lat"`
	Long      string `json:"lang"`
	Radius    string `json:"radius"`
	PlaceType string `json:"type"`
	Creator   string `json:"creator"`
}
