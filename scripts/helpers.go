package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       int
}

func NewConfig() (*Config, error) {
	srv := &Config{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     6379,
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	}

	if srv.Host == "" {
		return nil, fmt.Errorf("REDIS_HOST must be set")
	}

	if srv.Username == "" {
		return nil, fmt.Errorf("REDIS_USERNAME must be set")
	}

	if srv.Password == "" {
		return nil, fmt.Errorf("REDIS_PASSWORD must be set")
	}

	if port, err := strconv.Atoi(os.Getenv("REDIS_PORT")); err == nil {
		srv.Port = port
	} else {
		return nil, fmt.Errorf("REDIS_PORT must be an integer")
	}

	if db, err := strconv.Atoi(os.Getenv("REDIS_DB")); err == nil {
		srv.DB = db
	}

	return srv, nil
}

func NewRedisClient(conf *Config) (*redis.Client, error) {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: strings.TrimSpace(conf.Password),
		DB:       conf.DB,
	})

}

func ClientWithTimeout(ctx context.Context, client *redis.Client, timeout time.Duration) *redis.Client {
	return client.NewClientWithTimeout(ctx, timeout)
}

func GetOrSet(ctx context.Context, client *redis.Client, key string, value interface{}, ttl time.Duration) error {
	return client.Set(ctx, key, value, ttl).Err()
}

func Get(ctx context.Context, client *redis.Client, key string) (*string, error) {
	return client.Get(ctx, key).Result()
}