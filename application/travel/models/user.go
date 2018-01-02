package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

type MessageServiceProvider struct{}

var MessageService *MessageServiceProvider

type User struct {
	ID   int64
	Name string
	Sex  bool
}

func init() {
	orm.RegisterModel(new(User))
}

//Register
func (this *MessageServiceProvider) Register() (int64, error) {
	o := orm.NewOrm()
	o.Using("User")
	user := new(User)
	user.ID = 44
	user.Name = "Sansa Stack"
	user.Sex = true
	id, err := o.Insert(user)
	if err != nil {
		return 0, err
	}
	return id, err
}

//read user info
func (this *MessageServiceProvider) UserInfo() {
	info := User{ID: 44}
	fmt.Sprint(info.ID, info.Name, info.Sex)
	return
}
