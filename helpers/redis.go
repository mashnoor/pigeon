package helpers

import (
	"context"
	"github.com/mashnoor/pigeon/settings"
)

func SetRedis(key, value string) {
	client := settings.GetRedisClient()

	client.Set(context.Background(), key, value, 0)

}

func GetRedis(key string) (string, error) {

	client := settings.GetRedisClient()
	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil

}
