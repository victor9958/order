package routers

import (
	"github.com/astaxie/beego"
	"order/controllers"
)

func init() {
	//
	beego.Router("/v3/order/list", &controllers.GreensController{}, "get:Index")
	//beego.Router("/v3/order/ceshi", &controllers.GreensController{}, "get:Ceshi")
}
