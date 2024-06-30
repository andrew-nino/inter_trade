package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"

	"international_trade/config"
)

func ConnectRedis(cfg *config.Config) *redis.Client {

	redisUri := fmt.Sprintf("%s:%s", cfg.REDIS.Host, cfg.REDIS.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     redisUri,
		Password: cfg.REDIS.Password,
		DB:       0,
	})

	if _, redis_err := client.Ping().Result(); redis_err != nil {
		log.Fatal("Error: Unable to connect to Redis")
	}
	log.Println("Redis init was completed")

	return client
}
