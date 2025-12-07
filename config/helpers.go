package helpers

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	redis "github.com/go-redis/redis/v9"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedisClient(ctx context.Context, config RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewLogger(logLevel string) (*zap.Logger, error) {
	var logConfig zap.Config
	switch logLevel {
	case "debug":
		logConfig = zap.NewDevelopmentConfig()
	case "info":
		logConfig = zap.NewProductionConfig()
	case "error":
		logConfig = zap.NewProductionConfig()
	default:
		return nil, fmt.Errorf("unsupported log level: %s", logLevel)
	}
	logger, err := logConfig.Build()
	if err != nil {
		return nil, err
	}
	return logger, nil
}