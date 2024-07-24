package service

import (
	"encoding/json"
	"github.com/liangboceo/yuanboot/pkg/cache/redis"
	redisdb "github.com/liangboceo/yuanboot/pkg/datasources/redis"
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
	data := make(map[string]string)
	data["key"] = name
	dataString, _ := json.Marshal(data)
	_ = client.GetKVOps().Set(name, dataString, 0)
	json, _ := client.GetKVOps().GetString(name)
	return json
}
