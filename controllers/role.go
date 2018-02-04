package controllers

import(
	"time"
	"chain_admin.golang/models"
)

type RoleController struct{
	BaseController
}

func (self *RoleController) List(){
	roles, total := models.GetRoleList()
	self.Data["roles"] = roles
	self.Data["total"] = total
	if roles != nil && len(roles) > 0{
		self.Data["hasdata"] = true
	}else{
		self.Data["hasdata"] = false
	}
	self.display()
}

func (self *RoleController) RoleAdd(){
	self.modal_display()
}

func (self *RoleController) Add(){
	name := self.GetString("name")
	desc := self.GetString("desc")
	if len(name) < 1{
		self.ajaxMsg("角色名称不能为空", MSG_ERR)
	}
	role := new(models.Role)
	role.Name = name
	role.Desc = desc
	role.Status = 1
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	_, err := models.RoleAdd(role)
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("添加角色成功", MSG_OK)
}