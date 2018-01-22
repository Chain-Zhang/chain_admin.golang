package controllers

import (
	"strings"
	"github.com/astaxie/beego"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct{
	beego.Controller
	controllerName string
	actionName string
	userid int
	username string
	pageSize int
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
	self.Data["username"] = self.username
	self.Data["userid"] = self.userid
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