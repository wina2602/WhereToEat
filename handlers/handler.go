package handlers

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type handler struct {
	DB  *gorm.DB
	Rdb *redis.Client
}

func Resources(db *gorm.DB, rdb *redis.Client) handler {
	return handler{db, rdb}
}
