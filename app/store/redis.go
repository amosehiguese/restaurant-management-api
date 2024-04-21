package store

import (
	"fmt"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func RedisConn() (*redis.Client, error) {
	dbNumber, _ := strconv.Atoi(os.Getenv("REDIS_DB_NUMBER"))

	url := fmt.Sprintf("%s:%s",os.Getenv("REDIS_HOST"),os.Getenv("REDIS_PORT"))

	options := &redis.Options{
		Addr: url,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: dbNumber,
	}

	return redis.NewClient(options), nil

}