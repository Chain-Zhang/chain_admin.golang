package controllers

import (
	"strconv"
	"chain_admin.golang/util"
	"chain_admin.golang/models"
)

type LoginController struct{
	BaseController
}

func (self *LoginController)LoginIn(){
	self.TplName = "login/login.html"
}

func (self *LoginController)AjaxLogin(){
	username := self.GetString("username")
	password := self.GetString("password")
	password = util.Md5(password, false)
	user, err := models.GetUserByName(username)
	if err != nil{
		self.ajaxMsg("用户不存在", MSG_ERR)
	}
	if password != user.Password{
		self.ajaxMsg("密码不正确", MSG_ERR)
	}
	if user.Status != 1{
		self.ajaxMsg("该账户已被禁用", MSG_ERR)
	}
	authkey := util.Md5(self.getClientIp() + "|" + password, false)
	self.Ctx.SetCookie("auth", strconv.Itoa(user.Id) + "|" + authkey, 7 * 86400)
	self.ajaxMsg("登录成功", MSG_OK)
}