package redis_client

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

const redisPort = 6379

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient() *RedisClient {
	host := getHostFromEnvironment()
	address := fmt.Sprintf("%s:%d", host, redisPort)
	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr: address,
		}),
	}
}

func (redisClient *RedisClient) AddToList(jsonString string, listName string) error {
	fmt.Println("Received jsonstring: " + jsonString)
	fmt.Println("Received listname: " + listName)
	_, err := redisClient.client.RPush(context.Background(), listName, jsonString).Result()
	if err != nil {
		return err
	}

	return nil
}

func getHostFromEnvironment() string {
	if os.Getenv("environment") == "production" {
		fmt.Println("found environment=production")
		return "redis-service"
	} else {
		return "localhost"
	}
}
