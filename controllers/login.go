package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	//	"net/url"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"

}

func (this *LoginController) Put() {
	this.Ctx.WriteString(fmt.Sprint(this.Input))
	uname := this.Input().Get("uname")
	pwd := this.Input().Get("pwd")
	autoLogin := this.Input().Get("autoLogin") == "on"

	//正确与否的判断
	if beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autoLogin {
			maxAge = 1 << 31 / 1
		}
		//用cookie储存
		this.Ctx.SetCookie("uname", uname, maxAge, "/")
		this.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	//重定向
	this.Redirect("/", 301)
	return
}

//验证帐号
func checkAccount(ctx *context.Context) bool {
	ck, err := ctx.Request.Cookie("uname")
	if err != nil {
		return false
	}

	uname := ck.Value
	ck, err = ctx.Request.Cookie("pwd")
	if err != nil {
		return false
	}

	pwd := ck.Value

	return beego.AppConfig.String("uname") == uname &&
		beego.AppConfig.String("pwd") == pwd

}
