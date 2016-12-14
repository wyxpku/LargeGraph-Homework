package routers

import (
	"github.com/astaxie/beego"
	"github.com/user/LargeGraph-Homework/controllers"
)

func init() {
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/moment", &controllers.UserController{}, "GET:GetUserMoment")
	beego.Router("/user/following", &controllers.UserController{}, "GET:GetUserFollowing")
	beego.Router("/user/followed", &controllers.UserController{}, "GET:GetUserFollowed")

	beego.Router("/moment", &controllers.UserController{}, "GET:FriendMomentJson")
	beego.Router("/moment", &controllers.UserController{}, "POST:AddMoment")

	beego.Router("/follow", &controllers.UserController{}, "POST:AddFollow")
	beego.Router("/unfollow", &controllers.UserController{}, "POST:RemoveFollow")
}
