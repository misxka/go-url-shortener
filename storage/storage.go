package storage

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

func InitStorage() *StorageService {
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", port),
		Password: password,
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}\n", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, 24*time.Hour).Err()
	if err != nil {
		log.Fatalf("Failed saving key url | Error: %v\n", err)
	}
}

func GetOriginalUrl(shortUrl string) string {
	originalUrl, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		log.Fatalf("Failed retrieving original url | Error: %v\n", err)
	}
	return originalUrl
}
