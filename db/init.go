package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/riveryang/spider/models"
)

func InitDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/spider?charset=utf8", 3, 10)
	orm.RegisterModel(new(models.DmhyTopic), new(models.Bangumi), new(models.BangumiInfo), new(models.Character))

	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		beego.Error(err)
	}
}
