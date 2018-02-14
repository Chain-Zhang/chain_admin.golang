package controllers

import(
	"time"
	"strings"
	"strconv"
	"chain_admin.golang/models"

	"github.com/astaxie/beego/logs"
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
	ids := self.GetString("ids")
	if len(name) < 1{
		self.ajaxMsg("角色名称不能为空", MSG_ERR)
	}
	role := new(models.Role)
	role.Name = name
	role.Desc = desc
	role.Status = 1
	role.CreatedAt = time.Now()
	role.UpdatedAt = time.Now()
	roleId, err := models.RoleAdd(role)
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	idSlice := strings.Split(ids, ",")
	
	ra := new(models.RoleAuth)
	for _, v :=range idSlice{
		if len(v) == 0{
			continue
		}
		aid, _ := strconv.Atoi(v)
		ra.AuthId = aid
		ra.RoleId = roleId
		_, err = models.RoleAuthAdd(ra)
		if err != nil{
			logs.Error(err)
		}
	}
	self.ajaxMsg("添加角色成功", MSG_OK)
}

func (self *RoleController) RoleEdit(){
	id, _ := self.GetInt("id")
	role, _ := models.GetRoleById(id)
	self.Data["role"] = role

	ras, _ := models.GetRoleAuthList("role_id", id)
	authids := make([]int, 0)
	for _, v := range ras{
		authids = append(authids, v.AuthId)
	}
	self.Data["authids"] = authids
	self.modal_display()
}

func (self *RoleController)Edit(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	role, err := models.GetRoleById(id)
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	role.Name = self.GetString("name")
	role.Desc = self.GetString("desc")
	role.UpdatedAt = time.Now()
	
	err = role.Update()
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	ids := self.GetString("ids")
	idSlice := strings.Split(ids, ",")

	_, err = models.RoleAuthDelete(id)
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	ra := new(models.RoleAuth)
	for _, v :=range idSlice{
		if len(v) == 0{
			continue
		}
		aid, _ := strconv.Atoi(v)
		ra.AuthId = aid
		ra.RoleId = int64(id)
		_, err = models.RoleAuthAdd(ra)
		if err != nil{
			logs.Error(err)
		}
	}
	self.ajaxMsg("编辑角色成功", MSG_OK)
}

func (self *RoleController) Delete(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	role, err := models.GetRoleById(id)
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	_, err = role.Delete("id")
	if err != nil{
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("角色【" + role.Name +"】已删除", MSG_OK)
}