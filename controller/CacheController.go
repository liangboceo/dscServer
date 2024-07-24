package controller

import (
	"dscserver/service"
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

func (controller CacheController) GetTodoList() actionresult.IActionResult {
	return actionresult.Data{
		ContentType: "application/json; charset=utf-8",
		Data:        []byte(controller.cache.GetCache("dsc:database")),
	}
}
