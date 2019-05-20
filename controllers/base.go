package controllers

import (
	"github.com/astaxie/beego"
	"nim_163/server"
)

type NestPreparer interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
}

func (base *baseController) Prepare() {
	server.NewClient(beego.AppConfig.String("appKey"), beego.AppConfig.String("appSecret"))
	if app, ok := base.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}
