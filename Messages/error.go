package messages

var err Err

func UserNotFoundError() Err {
	setMsg("Invalid credentials.Please login again")
	return err
}

func UserNotCreated() Err {
	setMsg("Fail to create user")
	return err
}
func ChooseDifferentUname() Err {
	setMsg("Username already taken,Choose a differrent one")
	return err
}
func ServerError() Err {
	setMsg("Server connection failed.")
	return err
}
func setMsg(msg string) {
	err.Message = msg
}
