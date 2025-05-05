package storage

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type StorageService struct {
	redisClient *redis.Client
}

var ctx = context.Background()

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

	return &StorageService{redisClient: redisClient}
}

func (s *StorageService) SaveUrlMapping(shortUrl string, originalUrl string, userId string) error {
	err := s.redisClient.Set(ctx, shortUrl, originalUrl, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("Failed saving key url | Error: %v", err)
	}
	return nil
}

func (s *StorageService) GetOriginalUrl(shortUrl string) (string, error) {
	originalUrl, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key not found: %s", shortUrl)
	} else if err != nil {
		return "", fmt.Errorf("Failed retrieving original URL: %v", err)
	}
	return originalUrl, nil
}
