package main

import (
	"log"
	"where_to_eat/db"
	"where_to_eat/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	h := handlers.NewDb(DB)
	router := gin.Default()
	router.POST("/userCreate", h.Saveuser)
	router.POST("/login", h.Verifyuser)
	router.Run("localhost:4001")
	log.Println("Server started , listening to request at port 4001")
}
