package controller

import (
	"github.com/liangboceo/yuanboot/web/middlewares"
	"github.com/liangboceo/yuanboot/web/mvc"
)

type IndexController struct {
	mvc.ApiController // 必须继承
	log               *middlewares.Logger
}

func NewIndexController() *IndexController {
	indexController := &IndexController{}
	indexController.log = middlewares.NewLogger()
	return indexController
}

type RegisterRequest struct {
	mvc.RequestBody
	UserName string `param:"UserName"`
	Password string `param:"Password"`
}

func (controller IndexController) GetInfo() mvc.ApiResult {
	return controller.OK("ok")
}
