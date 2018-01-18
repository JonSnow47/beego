package models

import (
	"github.com/astaxie/beego/orm"

	"github.com/JonSnow47/beego/application/Fan/utility"
)

type AdminServiceProvider struct{}

var AdminServer *ArticleServiceProvider

type Admin struct {
	ID       int64  `orm:"column(id)";pk`
	Name     string `orm:"column(name)"	json:"name"`
	Password string `orm:"column(password)"	json:"password"`
	//State    bool   `orm:"column(state)"	json:"state"`
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
	Name string `orm="column(name)"`
	//Sex     string `orm="column(sex)"`
	//Age     int    `orm="column(age)"`
	//Address string `orm="column(address)"`
	//Email   string `orm="column(email)"`
}

func init() {
	orm.RegisterModel(new(Admin) /*, new(Profile)*/)
}

func (this *ArticleServiceProvider) Create(name string, pass string) error {
	o := orm.NewOrm()

	hash, err := utility.GenerateHash(pass)
	if err != nil {
		return err
	}
	password := string(hash)

	sql := "INSERT INTO Fan.admin(name,password) VALUES(?,?)"
	_, err = o.Raw(sql, name, password).Exec()

	return err
}

func (this *ArticleServiceProvider) Login(name string, password string) (bool, error) {
	o := orm.NewOrm()
	var pass string
	err := o.Raw("SELECT password FROM Fan.admin WHERE name=?", name).QueryRow(&pass)
	if err != nil {
		return false, err
	} else {
		if utility.CompareHash([]byte(pass), password) {
			return true, nil
		}
		return false, nil
	}
}

func (this *ArticleServiceProvider) AdminInfo(name string) (Info, error) {
	o := orm.NewOrm()

	var info Info
	err := o.Raw("SELECT * FROM Fan.admin WHERE name=? AND state=1", name).QueryRow(&info)
	return info, err
}
