package routers

import (
	"github.com/astaxie/beego"
	"github.com/user/LargeGraph-Homework/controllers"
)

func init() {
	beego.Router("/", &controllers.BaseController{})
	beego.Router("/json", &controllers.BaseController{}, "get:GetUserJson")
	beego.Router("/admin", &controllers.BaseController{}, "get:Admin")
	beego.Router("/admin/user", &controllers.BaseController{}, "get:AdminJson")
}
