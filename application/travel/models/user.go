package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type UserServiceProvider struct{}

var UserServer *UserServiceProvider

type User struct {
	ID   int64  `orm:"column(id);pk"	json:"id"`
	Name string `orm:"column(name)"	json:"name"`
	Pass string `orm:"column(pass)"	json:"pass"`
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
	//logs.Debug(info)
	id, err := o.Insert(&info)
	if err != nil {
		return 0, err
	}
	return id, err
}

//read user info
func (this *UserServiceProvider) ReadInfoByID(id int64) (User, error) {
	o := orm.NewOrm()
	o.Using("user")

	u := User{ID: id}
	//sql := ("SELECT * FROM user id = ?", u)
	err := o.Read(u)
	return u, err
}

//delete account
func (this *UserServiceProvider) Delete(content User) (int64, error) {
	o := orm.NewOrm()
	o.Using("user")

	id, err := o.Delete(&content)
	if err != nil {
		return 0, err
	}
	return id, err
}
