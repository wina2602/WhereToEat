package main

import (
	"log"
	"where_to_eat/db"
	"where_to_eat/handlers"
	redislocal "where_to_eat/redis"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	Rdb := redislocal.GetConnection()
	h := handlers.Resources(DB, Rdb)
	router := gin.Default()
	router.POST("/createUser", h.Saveuser)
	router.POST("/login", h.Verifyuser)
	router.POST("/createRoom", h.Createroom)
	router.POST("/joinRoom", h.JoinRoom)
	router.POST("/right", h.RightSwipes)
	router.GET("/getRooms", h.GetRooms)
	router.Run("localhost:4001")
	log.Println("Server started , listening to request at port 4001")
}
