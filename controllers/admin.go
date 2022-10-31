package controllers

import (
	"SimpleBlog/models"
	"SimpleBlog/util"
	"strings"
)

type AdminController struct {
	baseController
}

// 配置信息
func (c *AdminController) Config() {
	var result []*models.Config
	c.o.QueryTable(new(models.Config).TableName()).All(&result)
	options := make(map[string]string)
	mp := make(map[string]*models.Config)
	for _, v := range result {
		options[v.Name] = v.Value
		mp[v.Name] = v
	}
	if c.Ctx.Request.Method == "POST" {
		keys := []string{"url", "title", "keywords", "description", "email", "start", "qq"}
		for _, key := range keys {
			val := c.GetString(key)
			if _, ok := mp[key]; !ok {
				options[key] = val
				c.o.Insert(&models.Config{Name: key, Value: val})
			} else {
				opt := mp[key]
				if _, err := c.o.Update(&models.Config{Id: opt.Id, Name: opt.Name, Value: val}); err != nil {
					continue
				}
			}
		}
		c.History("设置数据成功", "")
	}
	c.Data["config"] = options
	c.TplName = c.controllerName + "/config.html"
}

// 后台用户登录
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		user := models.User{Username: username}
		c.o.Read(&user, "username")
		if user.Password == "" {
			c.History("账号不存在", "")
		}
		if util.Md5(password) != strings.Trim(user.Password, " ") {
			c.History("密码错误", "")
		}
		user.LoginCount = user.LoginCount + 1
		_, err := c.o.Update(&user)
		if err != nil {
			c.History("登录失败，写入信息失败", "")
		}
		c.History("登录成功", "/admin/main.html")
		c.SetSession("user", user)
	}
	c.TplName = c.controllerName + "/login.html"
}
