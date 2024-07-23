package controller

import (
	"github.com/liangboceo/yuanboot/web/context"
	"github.com/liangboceo/yuanboot/web/middlewares"
	"github.com/liangboceo/yuanboot/web/mvc"
)

type DemoController struct {
	mvc.ApiController // 必须继承
	log               *middlewares.Logger
}

func NewDemoController() *DemoController {
	demoController := &DemoController{}
	demoController.log = middlewares.NewLogger()
	return demoController
}

// -------------------------------------------------------------------------------
type RegisterRequest struct {
	mvc.RequestBody
	UserName string `param:"UserName"`
	Password string `param:"Password"`
}

// GET URL  http://localhost:8080/app/v1/demo/register?UserName=max&Password=123
func (controller DemoController) Register(ctx *context.HttpContext, request *RegisterRequest) mvc.ApiResult {
	defer func() {
		controller.log.ALogger.Info("hello")
	}()
	return mvc.ApiResult{Success: true, Message: "ok", Data: request}
}

// GET URL http://localhost:8080/app/v1/demo/getinfo
func (controller DemoController) GetInfo() mvc.ApiResult {
	return controller.OK("ok")
}
