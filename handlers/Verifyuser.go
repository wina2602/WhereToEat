package handlers

import (
	messages "where_to_eat/Messages"
	"where_to_eat/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) Verifyuser(c *gin.Context) {
	var userData models.Saverequest
	c.BindJSON(&userData)
	uname := userData.Username
	pwd := userData.Password

	var user models.Usercreds
	result := h.DB.Where("username =?", uname).First(&user)
	if result.Error != nil {
		c.IndentedJSON(401, messages.UserNotFoundError())

		return
	} else {
		pwd = pwd + user.Salt
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pwd)); err != nil {
			c.IndentedJSON(401, messages.UserNotFoundError())
			return

		} else {
			c.IndentedJSON(200, messages.SuccessMsg(user))
			return
		}
	}

}
