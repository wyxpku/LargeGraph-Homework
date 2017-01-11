package routers

import (
	"github.com/astaxie/beego"
	"github.com/user/LargeGraph-Homework/controllers"
)

func init() {
	beego.Router("/signup", &controllers.AccountController{}, "get:GetSignup")
	beego.Router("/login", &controllers.AccountController{}, "get:GetLogin")
	beego.Router("/account/signup", &controllers.AccountController{}, "post:SignUp")
	beego.Router("/account/login", &controllers.AccountController{}, "post:Login")
	beego.Router("/account/logout", &controllers.AccountController{}, "post:Logout")
}
