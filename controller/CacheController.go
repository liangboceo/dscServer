package controller

import (
	"dscserver/dto"
	"dscserver/service"
	"encoding/json"
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
	Url        string `param:"url" doc:"记录请求的url"`
}

// 缓存设置请求体
type CacheSetReq struct {
	mvc.RequestBody
	ServerName string `json:"serverName" doc:"服务名"`
	KeyName    string `json:"keyName" doc:"缓存md5Key"`
	Value      string `json:"value" doc:"缓存的值"`
	Expire     int    `json:"expire" doc:"过期时间"`
}

// GetFrontCache 获取前台服务缓存
func (controller CacheController) GetFrontCache(req *CacheReq) actionresult.IActionResult {
	key := fmt.Sprintf("%s:%s", req.ServerName, req.KeyName)
	value := controller.cache.GetCache(key)
	res, _ := json.Marshal(dto.Success(value))
	return actionresult.Data{
		ContentType: "application/json; charset=utf-8",
		Data:        res,
	}
}

func (controller CacheController) SetFrontCache(req *CacheSetReq) actionresult.IActionResult {
	key := fmt.Sprintf("%s:%s", req.ServerName, req.KeyName)
	value := controller.cache.SetCache(key, req.Value, req.Expire)
	var res []byte
	if !value {
		res, _ = json.Marshal(dto.Failure(value))
	} else {
		res, _ = json.Marshal(dto.Success(value))
	}
	return actionresult.Data{
		ContentType: "application/json; charset=utf-8",
		Data:        res,
	}

}
