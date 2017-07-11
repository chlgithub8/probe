package routers

import (
	"probe/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/GetJson", &controllers.MainController{}, "get:GetJson")
}
