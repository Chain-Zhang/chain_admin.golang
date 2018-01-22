package routers

import (
	"chain_admin.golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/login", &controllers.LoginController{},"*:LoginIn")
	beego.Router("/admin/index",&controllers.IndexController{},"*:Index")
	beego.Router("/admin/userlist", &controllers.UserController{},"*:UserList")
	beego.Router("/admin/useradd", &controllers.UserController{},"*:UserAdd")

	beego.Router("/service/admin/user_add", &controllers.UserController{},"*:AjaxAdd")
	beego.Router("/service/admin/user_status_change", &controllers.UserController{}, "*:UserStatusChange")
}
