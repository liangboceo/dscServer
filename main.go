package main

import (
	"dscserver/controller"
	"github.com/liangboceo/dependencyinjection"
	"github.com/liangboceo/yuanboot/abstractions"
	nacosconfig "github.com/liangboceo/yuanboot/pkg/configuration/nacos"
	"github.com/liangboceo/yuanboot/pkg/servicediscovery/nacos"
	"github.com/liangboceo/yuanboot/pkg/swagger"
	"github.com/liangboceo/yuanboot/web"
	"github.com/liangboceo/yuanboot/web/actionresult/extension"
	"github.com/liangboceo/yuanboot/web/context"
	"github.com/liangboceo/yuanboot/web/endpoints"
	"github.com/liangboceo/yuanboot/web/middlewares"
	"github.com/liangboceo/yuanboot/web/mvc"
	"github.com/liangboceo/yuanboot/web/router"
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
			app.UseMiddleware(middlewares.NewCORS())
			//web.UseMiddleware(middlewares.NewRequestTracker())
			app.UseStaticAssets()
			app.UseEndpoints(registerEndpointRouterConfig)
			app.UseMvc(func(builder *mvc.ControllerBuilder) {
				builder.AddViewsByConfig() //视图
				builder.EnableRouteAttributes()
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

func registerEndpointRouterConfig(rb router.IRouterBuilder) {
	endpoints.UseHealth(rb)
	endpoints.UseViz(rb)
	endpoints.UsePrometheus(rb)
	endpoints.UsePprof(rb)
	endpoints.UseReadiness(rb)
	endpoints.UseLiveness(rb)
	endpoints.UseJwt(rb)
	endpoints.UseRouteInfo(rb)
	endpoints.UseSwaggerDoc(rb,
		swagger.Info{
			Title:          "yuanboot 框架文档演示",
			Version:        "v1.0.0",
			Description:    "框架文档演示swagger文档 v1.0 [ #yuanboot](https://github.com/liangboceo/yuanboot).",
			TermsOfService: "https://dev.yuanboot.run",
			Contact: swagger.Contact{
				Email: "liangboceo@hotmail.com",
				Name:  "yuanboot",
			},
			License: swagger.License{
				Name: "MIT",
				Url:  "https://opensource.org/licenses/MIT",
			},
		},
		func(openapi *swagger.OpenApi) {
			openapi.AddSecurityBearerAuth()
		})

	rb.GET("/error", func(ctx *context.HttpContext) {
		panic("http get error")
	})

}
