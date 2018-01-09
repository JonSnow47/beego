package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminServiceProvider struct {}
var AdminServer *ArticleServiceProvider

type Admin struct{
	ID	int64
	Name	string
	Password	string
	Profile	Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

func init() {
	orm.RegisterModel(new(Admin),new(Profile))
}

func (this *ArticleServiceProvider) New()  {

}