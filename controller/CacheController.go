package controller

import (
	"github.com/liangboceo/yuanboot/web/middlewares"
	"github.com/liangboceo/yuanboot/web/mvc"
)

type CacheController struct {
	mvc.ApiController
	log *middlewares.Logger
}

func NewCacheController() *CacheController {
	cacheController := &CacheController{}
	cacheController.log = middlewares.NewLogger()
	return cacheController
}
