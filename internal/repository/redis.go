package repository

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() error {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	err := RedisClient.Ping(Ctx).Err()
	if err != nil {
		log.Fatalf("Redis connection failed: %v\n", err)
		return err
	}

	log.Println("Connected to Redis!")
	return nil
}
