package handlers

import (
	"context"
	messages "where_to_eat/Messages"
	"where_to_eat/models"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func (h handler) JoinRoom(c *gin.Context) {
	var joinRequest models.JoinroomRequest
	c.BindJSON(&joinRequest)
	ctx := context.Background()
	isPresent, data := checkInRedis(ctx, joinRequest, h.Rdb)
	if !isPresent {
		c.AbortWithStatusJSON(500, data.Message)
		return
	}
	c.IndentedJSON(200, data)
	return
}

func checkInRedis(ctx context.Context, data models.JoinroomRequest, Rdb *redis.Client) (bool, messages.JoinedRoomMsg) {
	key := "roomData:" + data.SID

	isPresent, _ := Rdb.Exists(ctx, key).Result()
	if isPresent == 0 {
		var err messages.JoinedRoomMsg
		err.Message = "Room not present or expired!!"
		return false, err
	}
	Rdata, _ := Rdb.HGetAll(ctx, key).Result()
	joinedRoomResponse := messages.JoinedRoomMsg{Message: "Welcome to Room!May you your FOODY place!!", Lat: Rdata["Lat"], Long: Rdata["Long"], Radius: Rdata["Radius"], Creator: Rdata["Creator"], PlaceType: Rdata["PlaceType"]}
	return true, joinedRoomResponse

}
