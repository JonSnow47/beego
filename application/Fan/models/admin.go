package models

import (
	"github.com/astaxie/beego/orm"
)

type AdminServiceProvider struct{}

var AdminServer *ArticleServiceProvider

type Admin struct {
	ID       int64    `orm:column(id);pk`
	Name     string   `orm:"column(name)"	json:"name"`
	Password string   `orm:"column(password)"	json:"password"`
	State    bool     `orm:"column(state)"	json:"state"`
	//Profile  *Profile `orm:"rel(one)"`
}
/*
type Profile struct {
	Sex     string `orm:"column(sex)"	json:"sex"`
	Age     int    `orm:"column(age)"	json:"age"`
	Address string `orm:"column(address)"	json:"address"`
	Email   string `orm:"column(email)"	json:"email"`
}
*/
type Info struct {
	Name    string `orm="column(name)"`
	//Sex     string `orm="column(sex)"`
	//Age     int    `orm="column(age)"`
	//Address string `orm="column(address)"`
	//Email   string `orm="column(email)"`
}

func init() {
	orm.RegisterModel(new(Admin)/*, new(Profile)*/)
}

func (this *ArticleServiceProvider) Create(admin Admin) error {
	o := orm.NewOrm()
	//o.Using("admin")

	sql := "INSERT INTO Fan.admin(name,password,state) VALUES(?,?,?)"
	values := []interface{}{admin.Name, admin.Password, 1}
	_, err := o.Raw(sql, values).Exec()

	return err
}

func (this *ArticleServiceProvider) Signin(name string, pass string) error {
	o := orm.NewOrm()

	err := o.Raw("SELECT pass FROM Fan.admin WHERE name=? LIMIT 1 LOCK IN SHARE MODE", name).QueryRow(&pass)

	return err
}

func (this *ArticleServiceProvider) Signout() {

}
func (this *ArticleServiceProvider) AdminInfo(name string) (Info, error) {
	o := orm.NewOrm()

	var info Info
	err := o.Raw("SELECT * FROM Fan.admin WHERE name=? AND state=1", name).QueryRow(&info)
	return info, err
}
