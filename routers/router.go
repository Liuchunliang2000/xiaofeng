package routers

import (
	"HellBeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/login", &controllers.MainController{})
}
