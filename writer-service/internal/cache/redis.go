package cache

import (
	"context"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()
var Client *redis.Client

func InitRedis() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := Client.Ping(Ctx).Result()
	if err != nil {
		fmt.Printf("failed to connect redis: %v\n", err)
	}

	fmt.Println("Connected to Redis")
}


func ResolveCacheKey(parts ...string) string {
	return strings.Join(parts, ":")
}
