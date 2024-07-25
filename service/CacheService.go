package service

import (
	"github.com/liangboceo/yuanboot/pkg/cache/redis"
	redisdb "github.com/liangboceo/yuanboot/pkg/datasources/redis"
	"time"
)

type CacheService struct {
	redisClient *redisdb.RedisDataSource
}

func NewCacheService(redis *redisdb.RedisDataSource) *CacheService {
	return &CacheService{redisClient: redis}
}

func (cacheService *CacheService) GetCache(name string) string {
	conn, _, _ := cacheService.redisClient.Open()
	client := conn.(redis.IClient)
	json, _ := client.GetKVOps().GetString(name)
	return json
}

func (cacheService *CacheService) SetCache(name string, value string, expire int) bool {
	conn, _, _ := cacheService.redisClient.Open()
	client := conn.(redis.IClient)
	err := client.GetKVOps().Set(name, value, time.Duration(expire))
	if err == nil {
		return true
	}
	return false
}
