package controller

import (
	"github.com/liangboceo/yuanboot/web/middlewares"
	"github.com/liangboceo/yuanboot/web/mvc"
)

type DataInterFaceController struct {
	mvc.ApiController `doc:"数据接口"`
	log               *middlewares.Logger
}

func NewDataInterFaceController() *DataInterFaceController {
	dataInterfaceController := &DataInterFaceController{}
	dataInterfaceController.log = middlewares.NewLogger()
	return dataInterfaceController
}
