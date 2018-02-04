package models

import(
	"github.com/astaxie/beego/orm"
	"time"
)

type Role struct{
	Id int
	Name string
	Desc string
	Status int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (self *Role)TableName() string{
	return "sys_role"
}

func RoleAdd(role *Role)(int64, error){
	return orm.NewOrm().Insert(role)
}

func GetRoleList(filters ...interface{})([]*Role, int64){
	roles := make([]*Role, 0)
	query := orm.NewOrm().QueryTable("sys_role")
	if len(filters) > 0{
		l := len(filters)
        for i := 0; i < l; i += 2{
			query = query.Filter(filters[i].(string), filters[i + 1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").All(&roles)
	return roles, total
}

func GetRoleById(id int)(*Role, error){
	role := new(Role)
	err := orm.NewOrm().QueryTable(role.TableName()).Filter("id", id).One(role)
	if err != nil{
		return nil, err
	}
	return role, nil
}

func (self *Role) Update(fields ...string)error{
	_, err := orm.NewOrm().Update(self, fields...)
	return err
}

func (self *Role) Delete(fields ...string)(int64, error){
	return orm.NewOrm().Delete(self, fields...)
}