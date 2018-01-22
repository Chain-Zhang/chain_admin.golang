package controllers

type LoginController struct{
	BaseController
}

func (self *LoginController)LoginIn(){
	self.TplName = "login/login.html"
}