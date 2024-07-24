package controller

import (
	"dscserver/service"
	"fmt"
	"github.com/liangboceo/yuanboot/abstractions"
	"github.com/liangboceo/yuanboot/web/actionresult"
	"github.com/liangboceo/yuanboot/web/context"
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
	Key        string `param:"key" doc:"缓存md5Key"`
}

// GetFrontCache 获取前台服务缓存
func (controller CacheController) GetFrontCache(ctx *context.HttpContext, req *CacheReq) actionresult.IActionResult {
	var config abstractions.IConfiguration
	_ = ctx.RequiredServices.GetService(&config)
	key := fmt.Sprintf("%s:%s:%s", config.GetString("yuanboot.application.name"), req.ServerName, req.Key)
	return actionresult.Data{
		ContentType: "application/json; charset=utf-8",
		Data:        []byte(controller.cache.GetCache(key)),
	}
}
