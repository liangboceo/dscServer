package main

import (
	"dscserver/controller"
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions"
	nacosconfig "github.com/liangboceo/yuanboot/pkg/configuration/nacos"
	"github.com/liangboceo/yuanboot/pkg/servicediscovery/nacos"
	"github.com/liangboceo/yuanboot/web"
	"github.com/liangboceo/yuanboot/web/actionresult/extension"
	"github.com/liangboceo/yuanboot/web/mvc"
)

func main() {
	CreateMVCBuilder().Build().Run()
}

// * Create the builder of Web host
func CreateMVCBuilder() *abstractions.HostBuilder {
	configuration := nacosconfig.RemoteConfig("bootstrap")
	return web.NewWebHostBuilder().
		UseConfiguration(configuration).
		Configure(func(app *web.ApplicationBuilder) {
			app.SetJsonSerializer(extension.CamelJson())
			app.UseMvc(func(builder *mvc.ControllerBuilder) {
				builder.AddViewsByConfig()                           //视图
				builder.AddController(controller.NewIndexController) // 注册mvc controller
				builder.AddController(controller.NewCacheController)
				builder.AddController(controller.NewDataInterFaceController)
			})
		}).
		ConfigureServices(func(serviceCollection *dependencyinjection.ServiceCollection) {
			// ioc
			nacos.UseServiceDiscovery(serviceCollection)
		})
}
