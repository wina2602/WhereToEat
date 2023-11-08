package handlers

import (
	"fmt"
	errormsg "where_to_eat/Messages"
	"where_to_eat/models"
	"where_to_eat/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Test() {
	fmt.Println("dfkjvndfj")
}
func (h handler) Saveuser(c *gin.Context) {
	uname := c.PostForm("uname")
	pwd := c.PostForm("pwd")
	salt, err := utils.Createsalt()

	if err != nil {
		c.IndentedJSON(500, errormsg.UserNotCreated())
		return
	}

	pwd = pwd + salt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pwd), 8)
	if err != nil {
		c.IndentedJSON(500, errormsg.UserNotCreated())
		return
	}
	user := models.Usercreds{Username: uname, Salt: salt, Password: string(hashedPassword)}
	result := h.DB.Create(&user)
	if result.Error != nil {
		c.IndentedJSON(409, errormsg.ChooseDifferentUname())
		return
	}
	c.IndentedJSON(200, errormsg.SuccessMsg(user))
	return
}
