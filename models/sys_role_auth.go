package models

import(
	"github.com/astaxie/beego/orm"
)

type RoleAuth struct{
	RoleId int64
	AuthId int `orm:"pk"`
}

func (self *RoleAuth)TableName()string{
	return "sys_role_auth"
}

func RoleAuthAdd(ra *RoleAuth)(int64, error){
	return orm.NewOrm().Insert(ra)
}

func (self *RoleAuth)Delete(fields ...string)(int64, error){
	return orm.NewOrm().Delete(self, fields...)
}

func RoleAuthDelete(roleid int)(int64, error){
	return orm.NewOrm().QueryTable("sys_role_auth").Filter("role_id", roleid).Delete()
}

func GetRoleAuthList(filters ...interface{})([]*RoleAuth, int64){
	ras := make([]*RoleAuth, 0)
	query := orm.NewOrm().QueryTable("sys_role_auth")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("auth_id").All(&ras)
	return ras, total
}