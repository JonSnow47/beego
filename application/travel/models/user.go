package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"fmt"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type UserServiceProvider struct{}

var UserServer *UserServiceProvider

type User struct {
	ID   int64  `orm:"column(id);pk"`//`json:"id" orm:"column(id);pk"`
	Name string `json:"name" orm:"column(name)"`
}

func init() {
	orm.RegisterModel(new(User))
}
func TableName() string {
	return "user"
}

//Register
func (this *UserServiceProvider) Register(info User) (int64, error) {
	o := orm.NewOrm()
	o.Using("user") //use databases:travel
	//user := new(Travel)
	logs.Debug(info)
	id, err := o.Insert(&info)
	if err != nil {
		return 0, err
	}
	return id, err
}

//read user info
func (this *UserServiceProvider) ReadInfobyID(id int64) error{
	o := orm.NewOrm()
	o.Using("user")
	u := &User{ID: id}
	// 	err = Ormer.Read(u)
	err := o.Read(u)
	return err
}
