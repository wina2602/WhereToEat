package handlers

import (
	"context"
	"strconv"
	"time"
	messages "where_to_eat/Messages"
	"where_to_eat/models"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func (h handler) Createroom(c *gin.Context) {
	ctx := context.Background()
	var request models.Request
	c.BindJSON(&request)
	radius := request.Radius
	placeType := request.PlaceType
	id := request.Id
	uname := request.Username
	lat := request.Lat
	Long := request.Long
	if len(uname) == 0 || len(placeType) == 0 || lat == 0 || Long == 0 || radius == 0 {
		c.AbortWithStatus(422)
		return
	}
	time := time.Now().Unix()
	sid := strconv.Itoa(id) + strconv.Itoa(int(time))
	room := models.Room{SID: sid, Lat: lat, Long: Long, PlaceType: placeType, Radius: radius, Creator: uname}
	result := storeInRedis(ctx, room, h.Rdb)
	if !result {
		err := messages.ServerError()
		c.AbortWithStatusJSON(500, err)
		return
	}
	SuccessMsg := messages.RoomCreatedMsg{SID: sid, Lat: lat, Long: Long, Radius: radius, Creator: uname, Message: "Room created successfully"}
	c.IndentedJSON(200, SuccessMsg)
	return
}

func storeInRedis(ctx context.Context, room models.Room, Rdb *redis.Client) bool {
	mapName := "roomData:" + room.SID
	hashMapData := map[string]interface{}{
		"SID":       room.SID,
		"Creator":   room.Creator,
		"Long":      room.Long,
		"Lat":       room.Lat,
		"Radius":    room.Radius,
		"PlaceType": room.PlaceType,
		// Add more field-value pairs as needed
	}
	_, err := Rdb.HSet(ctx, mapName, hashMapData).Result()
	ext := 86400 * time.Second
	Rdb.Expire(ctx, mapName, ext)
	if err != nil {
		return false
	}
	return true
}
