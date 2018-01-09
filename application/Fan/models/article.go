package models

import (
	"github.com/astaxie/beego/orm"
)

type ArticleServiceProvider struct{}

var ArticleServer *ArticleServiceProvider

type Article struct {
	ID       int64  `orm:"column(id);pk"	json:"id"`
	Title    string `orm:"column(title)"	json:"title"`
	Author   string `orm:"column(author)"	json:"author"`
	Class    string `orm:"column(class)"	json:"class""`
	Content  string `orm:"column(content)"	json:"content""`
	Views    int64  `orm:"column(views)"	json:"views""`
	State    bool   `orm:"column(state)"	json:"state"`
	//DateTime orm.DateTimeField
}

type Display struct {
	Title	string
	Class	string
	Content	string
}

func Init() {
	orm.RegisterModel(new(Article))
}
func (this *ArticleServiceProvider) New(article Article) error {
	o := orm.NewOrm()
	raw := o.Raw("INSERT INTO article (title,author,class,content) VALUES(?,?,?,?)", article.Title, article.Author, article.Class, article.Content)
	_, err := raw.Exec()
	return err
}

func (this *ArticleServiceProvider) Updata(article Article) error {
	o := orm.NewOrm()
	sql := "UPDATA article SET title=? AND class=? AND content=?"
	values := []interface{}{article.Title,article.Class,article.Content}
	_,err := o.Raw(sql,values).Exec()
	return err
}

func (this *ArticleServiceProvider) Read(id int64) error {
	o := orm.NewOrm()
	_,err := o.Raw("SELECT * FROM article WHERE id=?",id).Exec()
	return err

}