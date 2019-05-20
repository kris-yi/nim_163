package main

import (
	_ "nim_163/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.SetLogger("file", `{"filename":"logs/log.log"}`)
	beego.SetLogFuncCall(true)
	beego.SetLevel(beego.LevelInformational)
	beego.Run()
}
