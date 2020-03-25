package routers

import (

	"net/http"
	"net/http/pprof"
	"github.com/astaxie/beego"
	"practice/web/controllers"
)

func init() {
	beego.NewControllerRegister()
	beego.NewTree()
	beego.NewApp()
	app := beego.Router("/", &controllers.MainController{})
	handler := app.Handlers
	handler.Handler("/debug/pprof", http.HandlerFunc(pprof.Index))
	handler.Handler("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
	handler.Handler("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
	handler.AddAutoPrefix("/v1", &controllers.FollowController{})

}
