package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	client *redis.Client
)

func InitClient() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		PoolSize: 100,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = client.Ping(ctx).Result()
	return err
}

func HashSet(key, field string, data interface{}) error {
	err := client.HSet(context.Background(), key, field, data).Err()
	if err != nil {
		return fmt.Errorf("Redis HSet Error: %w ", err)
	}
	return err
}

func HashGet(key, field string) (string, error) {
	result := ""
	val, err := client.HGet(context.Background(), key, field).Result()
	if err == redis.Nil {
		return result, nil
	} else if err != nil {
		return result, err
	}
	return val, nil
}
