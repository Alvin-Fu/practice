package routers

import (
	"practice/websvr/controllers"
	"github.com/astaxie/beego"
	)

func RegRouter() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/debug/pprof/", &controllers.ProfController{})
	beego.Router("/debug/pprof/:app([\\w]+)", &controllers.ProfController{})
}
