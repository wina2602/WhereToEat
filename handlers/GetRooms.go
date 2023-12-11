package handlers

import (
	"context"
	"encoding/json"
	"strconv"
	messages "where_to_eat/Messages"

	"github.com/gin-gonic/gin"
)

func (h handler) GetRooms(c *gin.Context) {
	id := c.Query("uid")
	mapName := "roomData:" + id
	ctx := context.Background()
	keys, _, err := h.Rdb.Scan(ctx, 0, mapName+"*", 1000).Result()
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	var rooms []messages.JoinedRoomMsg
	for ind, key := range keys {
		var room messages.JoinedRoomMsg
		Rdata, _ := h.Rdb.HGetAll(ctx, key).Result()
		json.Marshal(Rdata)
		room.Lat = Rdata["Lat"]
		room.Long = Rdata["Long"]
		room.PlaceType = Rdata["PlaceType"]
		room.Radius = Rdata["Radius"]
		room.Message = "Room : " + strconv.Itoa(ind)
		rooms = append(rooms, room)
	}
	c.IndentedJSON(200, rooms)
	return
}
