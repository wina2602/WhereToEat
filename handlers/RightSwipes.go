package handlers

import (
	"context"
	"where_to_eat/models"

	"github.com/gin-gonic/gin"
)

func (h handler) RightSwipes(c *gin.Context) {
	ctx := context.Background()
	var swipeData models.Swiperequest
	c.BindJSON(&swipeData)
	sid := swipeData.SID
	rid := swipeData.RID
	mapName := "restaurantData:" + sid
	if err := h.Rdb.HGetAll(ctx, mapName); err == nil {
		err := h.Rdb.HSet(ctx, mapName, rid, 1).Err()
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
	} else {
		ridPresent, err := h.Rdb.HExists(ctx, mapName, rid).Result()
		if err != nil {
			c.AbortWithStatus(500)
			return
		}
		if ridPresent {
			h.Rdb.HIncrBy(ctx, mapName, rid, 1)

		} else {
			h.Rdb.HSet(ctx, mapName, rid, 1)
		}
	}
	c.Status(200)
	return
}
