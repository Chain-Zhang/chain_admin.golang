package routers

import (
	"chain_admin.golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{},"*:LoginIn")
	beego.Router("/admin/index",&controllers.IndexController{},"*:Index")
	beego.Router("/admin/userlist", &controllers.UserController{},"*:UserList")
	beego.Router("/admin/useradd", &controllers.UserController{},"*:UserAdd")
	beego.Router("/admin/useredit", &controllers.UserController{},"*:UserEdit")


	beego.Router("/admin/auth", &controllers.AuthController{},"*:List")

	beego.Router("/service/login", &controllers.LoginController{},"*:AjaxLogin")
	beego.Router("/service/loginout", &controllers.LoginController{},"*:LoginOut")
	beego.Router("/service/admin/user_add", &controllers.UserController{},"*:AjaxAdd")
	beego.Router("/service/admin/user_edit", &controllers.UserController{},"*:AjaxEdit")
	beego.Router("/service/admin/user_status_change", &controllers.UserController{}, "*:UserStatusChange")
	beego.Router("/service/admin/user_del", &controllers.UserController{}, "*:UserDel")
}
