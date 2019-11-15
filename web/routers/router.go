package routers

import (
	"github.com/astaxie/beego"
	"practice/web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
