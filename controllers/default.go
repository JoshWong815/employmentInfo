package controllers

import (
	"employmentInfo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)
var Sessionid interface{}
type MainController struct {
	beego.Controller
}
func (c *MainController) URLMapping() {
	c.Mapping("Index", c.Index)
	c.Mapping("Login", c.Login)
	c.Mapping("LoginTest", c.LoginTest)
	c.Mapping("Logout", c.Logout)

}
func (c *MainController)SessionTest() {
	c.Data["id"] = c.GetSession("id")
	if c.Data["id"] == nil {
		c.Redirect("/login", 302)
	}
	fmt.Println("id:",c.Data["id"])
	fmt.Println("ceshi")
}
func(c *MainController) GetSessionNum() int {
	return beego.GlobalSessions.GetActiveSession()
}
func (c *MainController) Index() {
	c.Data["id"]=c.GetSession("id")
	c.TplName = "index.html"
	if c.Data["id"]==nil{
		c.Redirect("/login",302)
	}
	i:=c.GetSessionNum()
	fmt.Println("当前活跃的session数：",i)
}
func (c *MainController) Login() {
	c.TplName = "login.html"
}
func (c *MainController) LoginTest() {
	//c.TplName="login.html"
	id := c.Input().Get("id")
	password := c.Input().Get("password")
		o := orm.NewOrm()
		var student models.Student
		qs := o.QueryTable(student)
		err1 := qs.Filter("id", id).Filter("password", password).One(&student)

		if err1 == nil {
			//fmt.Println(user.name,user.Password)
			c.SetSession("id",id)
			c.Data["id"]=c.GetSession("id")
			//Sessionid=c.CruSession.SessionID()
			c.Redirect("/index", 302)

		} else if err1 == orm.ErrNoRows {
			str := "用户名或密码输入错误!"
			c.Data["info"] = str
			c.TplName = "login.html"
			//fmt.Println("用户名或密码输入错误")
			//fmt.Println(user.Age)

		} else if err1 == orm.ErrMissPK {
			fmt.Println("找不到主键")
			c.Redirect("/login", 302)
		}

	}

func (c *MainController) Logout(){
	c.DelSession("id")
	c.Redirect("/",302)
}

//func (c *MainController)SessionTest(){
//	if c.Data["id"]==nil{
//		c.Redirect("/login",302)
//	}
//}

