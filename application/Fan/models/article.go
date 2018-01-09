package models

import (
	"github.com/astaxie/beego/orm"
)

type ArticleServiceProvider struct {}
var ArticleServer *ArticleServiceProvider

type Article struct {
	ID   int64	`orm:"column(id);pk"	json:"id"`
	Title	string	`orm:"column(title)"	json:"title"`
	Auther	string	`orm:"column(auther)"	json:"auther"`
	Class	string	`orm:"column(class)"	json:"class""`
	Content	string	`orm:"column(content)"	json:"content""`
	Views	int64	`orm:"column(views)"	json:"views""`
	DateTime	orm.DateTimeField
}

func Init()  {
	orm.RegisterModel(new(Article))
}
func (n *ArticleServiceProvider) New()  {

}