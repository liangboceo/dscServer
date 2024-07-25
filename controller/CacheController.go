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
	mvc.ApiController `doc:"前台缓存"`
	cache             *service.CacheService
}

func NewCacheController(cacheService *service.CacheService) *CacheController {
	return &CacheController{cache: cacheService}
}

// CacheReq 缓存请求体
type CacheReq struct {
	mvc.RequestGET `route:"/cache/getFrontCache" doc:"获取前台缓存"`
	ServerName     string `uri:"serverName" doc:"服务名"`
	KeyName        string `uri:"keyName" doc:"缓存md5Key"`
	Url            string `uri:"url" doc:"记录请求的url"`
}

// CacheSetReq 缓存设置请求体
type CacheSetReq struct {
	mvc.RequestBody `route:"/cache/setFrontCache" doc:"设置前台缓存"`
	ServerName      string `json:"serverName" doc:"服务名"`
	KeyName         string `json:"keyName" doc:"缓存md5Key"`
	Value           string `json:"value" doc:"缓存的值"`
	Expire          int    `json:"expire" doc:"过期时间"`
}

// GetFrontCache 获取前台服务缓存
func (controller CacheController) GetFrontCache(req *CacheReq) actionresult.IActionResult {
	var res []byte
	if req.ServerName == "" || req.KeyName == "" {
		res, _ = json.Marshal(dto.FailureMessage("", "服务名/缓存md5Key不能为空"))
		return actionresult.Data{
			ContentType: "application/json; charset=utf-8",
			Data:        res,
		}
	}
	key := fmt.Sprintf("%s:%s", req.ServerName, req.KeyName)
	value := controller.cache.GetCache(key)
	res, _ = json.Marshal(dto.Success(value))
	return actionresult.Data{
		ContentType: "application/json; charset=utf-8",
		Data:        res,
	}
}

// SetFrontCache  设置缓存前台服务
func (controller CacheController) SetFrontCache(req *CacheSetReq) actionresult.IActionResult {
	var res []byte
	if req.ServerName == "" || req.KeyName == "" {
		res, _ = json.Marshal(dto.FailureMessage("", "服务名/缓存md5Key不能为空"))
		return actionresult.Data{
			ContentType: "application/json; charset=utf-8",
			Data:        res,
		}
	}
	key := fmt.Sprintf("%s:%s", req.ServerName, req.KeyName)
	value := controller.cache.SetCache(key, req.Value, req.Expire)
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
