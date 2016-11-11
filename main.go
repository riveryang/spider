package main

import (
	_ "github.com/riveryang/spider/routers"
	"github.com/astaxie/beego"
	"github.com/riveryang/spider/db"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}

func init() {
	beego.SetLogger("file", `{"filename":"logs/spider.log"}`)
	beego.SetLevel(beego.LevelDebug)

	db.InitDB()
}