package models

import (
	"github.com/beego/beego"
	"github.com/beego/beego/v2/client/orm"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "24574"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Asia%2FShanghai"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User), new(Category), new(Post), new(Config), new(Comment))
}

func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
