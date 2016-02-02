package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"web_bee/controllers"
	"web_bee/models"
)

func init() {
	models.RegisterDB()

}
func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Run()

}
