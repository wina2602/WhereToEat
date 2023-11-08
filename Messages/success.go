package messages

import "where_to_eat/models"

var success_msg UserCreatedMsg

func SuccessMsg(user models.Usercreds) UserCreatedMsg {
	success_msg.Username = user.Username
	success_msg.ID = user.ID
	success_msg.Message = "Login Successful"
	return success_msg
}
