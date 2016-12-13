package routers

import (
	"github.com/astaxie/beego"
	"github.com/user/LargeGraph-Homework/controllers"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/admin", &controllers.BaseController{}, "get:Admin")
}
