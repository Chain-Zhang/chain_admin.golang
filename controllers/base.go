package controllers

import (
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"chain_admin.golang/util"
	"chain_admin.golang/models"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct{
	beego.Controller
	controllerName string
	actionName string
	pageSize int
	login_userId int
	login_user *models.User
}

func (self *BaseController) Prepare(){
	self.pageSize = 10
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["appversion"] = beego.AppConfig.String("appversion")
	self.Data["appname"] = beego.AppConfig.String("appname")
	self.Data["cur_controller"] = self.controllerName
	self.Data["cur_action"] = self.actionName

	if self.auth(){
		self.Data["login_nickname"] = self.login_user.Nickname
	}
}

func (self *BaseController)auth() bool{
	auth := self.Ctx.GetCookie("auth")
	arr := strings.Split(auth, "|")
	self.login_userId = 0
	if len(arr) == 2{
		idStr, authkey := arr[0], arr[1]
		self.login_userId, _ = strconv.Atoi(idStr)
		if self.login_userId > 0{
			user, _ := models.GetUserById(self.login_userId)
			if authkey == util.Md5(self.getClientIp() + "|" + user.Password, false){
				self.login_user = user
				return true
			}
		}
	}
	if self.login_userId < 1 && (self.controllerName != "login" && self.actionName != "loginin"){
		self.redirect("/login")
	}
	return false
}

func (self *BaseController) display(tpl ...string){
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/layout.html"
	self.TplName = tplname
}

func (self *BaseController) modal_display(tpl ...string){
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/modal_layout.html"
	self.TplName = tplname
}

func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (self *BaseController) redirect(url string) {
	logs.Info("url: ",url)
	self.Redirect(url, 302)
	self.StopRun()
}
