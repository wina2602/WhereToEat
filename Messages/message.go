package messages

type Err struct {
	Message string `json:"message"`
}
type UserCreatedMsg struct {
	ID       int    `json:"id"`
	Username string `json:"uname"`
	Message  string `json:"message"`
}
