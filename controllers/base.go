package controllers

import (
	"SimpleBlog/models"
	"strings"

	"github.com/beego/beego"
	"github.com/beego/beego/v2/client/orm"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}

// 控制器是admin并且不是登录操作，获取到的会员User为空，提示没有登录
// 控制器是bolg读取动作配置,读取数据库中的全部配置
func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()
	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)
	p.o = orm.NewOrm()
	//admin的非登录操作
	if strings.ToLower(p.controllerName) == "admin" && strings.ToLower(p.actionName) != "login" {
		if p.GetSession("user") == nil {
			p.History("未登录", "/admin/login")
		}
	}

	//初始化页面元素
	if strings.ToLower(p.controllerName) == "blog" {
		p.Data["actionName"] = strings.ToLower(actionName)
		var result []*models.Config
		p.o.QueryTable(new(models.Config).TableName()).All(&result)
		configs := make(map[string]string)
		for _, v := range result {
			configs[v.Name] = v.Value
		}
		p.Data["config"] = configs
	}
}

func (p *baseController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url, 302)
	}
}
