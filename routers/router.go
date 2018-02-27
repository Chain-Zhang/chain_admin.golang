package routers

import (
	"chain_admin.golang/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &controllers.LoginController{},"get:LoginIn")
	beego.Router("/admin/index",&controllers.IndexController{},"get:Index")
	beego.Router("/admin/userlist", &controllers.UserController{},"get:UserList")
	beego.Router("/admin/useradd", &controllers.UserController{},"get:UserAdd")
	beego.Router("/admin/useredit", &controllers.UserController{},"get:UserEdit")

	beego.Router("/admin/auth", &controllers.AuthController{},"get:List")

	beego.Router("/admin/role", &controllers.RoleController{},"get:List")
	beego.Router("/admin/roleadd", &controllers.RoleController{},"get:RoleAdd")
	beego.Router("/admin/roleedit", &controllers.RoleController{},"get:RoleEdit")


	beego.Router("/error/403", &controllers.ErrorController{},"get:Error403")
	beego.Router("/error/404", &controllers.ErrorController{},"get:Error404")




	beego.Router("/service/login", &controllers.LoginController{},"post:AjaxLogin")
	beego.Router("/service/loginout", &controllers.LoginController{},"post:LoginOut")
	beego.Router("/service/admin/user_add", &controllers.UserController{},"post:AjaxAdd")
	beego.Router("/service/admin/user_edit", &controllers.UserController{},"post:AjaxEdit")
	beego.Router("/service/admin/user_status_change", &controllers.UserController{}, "post:UserStatusChange")
	beego.Router("/service/admin/user_del", &controllers.UserController{}, "post:UserDel")

	beego.Router("/service/admin/auth_add", &controllers.AuthController{}, "post:Add")
	beego.Router("/service/admin/auth_getlist", &controllers.AuthController{}, "get:GetList")
	beego.Router("/service/admin/auth_getone", &controllers.AuthController{}, "get:GetOne")
	beego.Router("/service/admin/auth_edit", &controllers.AuthController{}, "post:Edit")
	beego.Router("/service/admin/auth_del", &controllers.AuthController{}, "post:Delete")

	beego.Router("/service/admin/role_add", &controllers.RoleController{}, "post:Add")
	beego.Router("/service/admin/role_edit", &controllers.RoleController{}, "post:Edit")
	beego.Router("/service/admin/role_del", &controllers.RoleController{}, "post:Delete")
}
