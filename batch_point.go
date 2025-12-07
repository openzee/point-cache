package flow

import (
	"github.com/redis/go-redis/v9"
)

type BatchPoint []*Point

func (BatchPoint) PushRedis(redisClient *redis.Client) error {

	pipe := redisClient.Pipeline()

}
