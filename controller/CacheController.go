package controller

import (
	"dscserver/service"
	"fmt"
	"github.com/liangboceo/yuanboot/web/actionresult"
	"github.com/liangboceo/yuanboot/web/mvc"
)

type CacheController struct {
	mvc.ApiController
	cache *service.CacheService
}

func NewCacheController(cacheService *service.CacheService) *CacheController {
	return &CacheController{cache: cacheService}
}

// 缓存请求体
type CacheReq struct {
	mvc.RequestBody
	ServerName string `param:"serverName" doc:"服务名"`
	KeyName    string `param:"keyName" doc:"缓存md5Key"`
}

// GetFrontCache 获取前台服务缓存
func (controller CacheController) GetFrontCache(req *CacheReq) actionresult.IActionResult {
	key := fmt.Sprintf("%s:%s", req.ServerName, req.KeyName)
	return actionresult.Data{
		ContentType: "application/json; charset=utf-8",
		Data:        []byte(controller.cache.GetCache(key)),
	}
}
