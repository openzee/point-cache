package flow

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

type CacheService struct {
	redisClient *redis.Client
	ctx         context.Context
}

func CreateCacheService(redis_url string) (*CacheService, error) {

	redis_opt, err := redis.ParseURL(redis_url)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	redisClient := redis.NewClient(redis_opt)

	for {
		if err := redisClient.Ping(ctx).Err(); err != nil {
			log.Errorf("Redis 连接失败: %v", err)
			time.Sleep(time.Second)
			continue
		}
		break
	}

	log.Infof("Redis 连接成功: %s", redis_url)

	return &CacheService{
		redisClient: redisClient,
		ctx:         ctx,
	}, nil
}

func (obj *CacheService) Cancel() {
	obj.ctx.Done()
}

func (obj *CacheService) RegisterChannel(c chan BatchPoint) {

	go func() {
		select {
		case <-obj.ctx.Done():
			log.Error("Stop Service")
			return
		case batch := <-c:
			batch.PushRedis(obj.redisClient)
		}
	}()
}
