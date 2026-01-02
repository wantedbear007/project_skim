package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/wantedbear007/project_skim/internal/env"
	"go.uber.org/zap"
)

func InitCache() (*redis.Client,error) {
	zap.L().Info("Init redis service")

	opt, err := redis.ParseURL(env.REDIS_URL.GetValue())

	if err != nil {
		zap.L().Error("Error initializing cache server ", 
		zap.Error(err),
	)

		return nil, err
	}

	client := redis.NewClient(opt)

	// if error := client.Ping(); error {

	// }

	zap.L().Info("Redis client initilized")

	return client, nil
}

