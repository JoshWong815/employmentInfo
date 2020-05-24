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
	c.Mapping("GetMainEchartOfCompanyCity", c.GetMainEchartOfCompanyCity)
	c.Mapping("GetMainEchartOfCompanyType", c.GetMainEchartOfCompanyType)


}
type CitysOfCompany struct{
	Value int `json:"value"`
	Name string `json:"name"`
}

type TypeOfCompany struct{
	Value int `json:"value"`
	Name string `json:"name"`
}
//返回首页用来统计学生签约单位性质的分布的饼状图的数据
func (c *MainController) GetMainEchartOfCompanyType(){

	e:=models.EveryStudentNewestEmployment()
	fmt.Println(e)
	var m []*TypeOfCompany
	for i:=range e{
		if e[i].Operation=="签约"{
			var typeOfCompany *string
			sql:="select type from company where name='"+e[i].Cname+"'"
			o:=orm.NewOrm()
			err:=o.Raw(sql).QueryRow(&typeOfCompany)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println(sql)
			var a TypeOfCompany
			a.Value++
			a.Name=*typeOfCompany
			fmt.Println("a:",a)
			m=append(m,&a)


		}

	}
	fmt.Println("m:",m)
	for i:=range m{
		fmt.Println(m[i])
	}
	map1:=make(map[string]int)
	for i:=range m{
		map1[m[i].Name]++
	}
	fmt.Println(map1)
	var resp []*TypeOfCompany
	for i:=range map1{
		var a TypeOfCompany
		a.Name=i
		a.Value=map1[i]
		resp=append(resp,&a)
	}
	c.Data["json"]=resp
	c.ServeJSON()

}
//返回首页用来统计学生签约城市的分布的饼状图的数据
func (c *MainController) GetMainEchartOfCompanyCity(){

	e:=models.EveryStudentNewestEmployment()
	fmt.Println(e)
	var m []*CitysOfCompany
	for i:=range e{
		if e[i].Operation=="签约"{
			var city *string
			sql:="select city from company where name='"+e[i].Cname+"'"
			o:=orm.NewOrm()
			err:=o.Raw(sql).QueryRow(&city)
			if err!=nil{
				fmt.Println(err)
			}
			fmt.Println(sql)
			var a CitysOfCompany
			a.Value++
			a.Name=*city
			fmt.Println("a:",a)
			m=append(m,&a)


		}

	}
	fmt.Println("m:",m)
	for i:=range m{
		fmt.Println(m[i])
	}
	map1:=make(map[string]int)
	for i:=range m{
		map1[m[i].Name]++
	}
	fmt.Println(map1)
    var resp []*CitysOfCompany
	for i:=range map1{
		var a CitysOfCompany
		a.Name=i
		a.Value=map1[i]
		resp=append(resp,&a)
	}
	c.Data["json"]=resp
	c.ServeJSON()
}


func (c *MainController) SessionTest() {
	c.Data["id"] = c.GetSession("id")
	if c.Data["id"] == nil {
		c.Redirect("/login", 302)
	}
	i := c.GetSessionNum()
	fmt.Println("当前活跃的session数：", i)

}
func (c *MainController) GetSessionNum() int {
	return beego.GlobalSessions.GetActiveSession()
}
func (c *MainController) Index() {
	c.Data["id"] = c.GetSession("id")
	c.TplName = "index.html"
	if c.Data["id"] == nil {
		c.Redirect("/login", 302)
	}
	i := c.GetSessionNum()
	fmt.Println("当前活跃的session数：", i)
}
func (c *MainController) Login() {
	c.TplName = "login.html"
}
func (c *MainController) LoginTest() {
	//c.TplName="login.html"
	id := c.Input().Get("id")
	password := c.Input().Get("password")

	var maps1 []orm.Params
	var maps2 []orm.Params
	o := orm.NewOrm()
	//var student models.Student
	//qs := o.QueryTable(student)

	sql1 := "select * from admin where id='" + id + "' and password='" + password + "'"
	fmt.Println(sql1)
	res1, err1 := o.Raw(sql1).Values(&maps1)

	sql2 := "select * from student where id='" + id + "' and password='" + password + "'"
	fmt.Println(sql2)
	res2, err2 := o.Raw(sql2).Values(&maps2)

	fmt.Println("res1的值为：", res1)
	fmt.Println("res2的值为：", res2)

	if err1 != nil {
		fmt.Println(err1)
	}
	if err2 != nil {
		fmt.Println(err2)
	}
	//err1 := qs.Filter("id", id).Filter("password", password).One(&student)
	//如果是管理员用户
	if res1 == 1 && res2 == 0 {
		//fmt.Println(user.name,user.Password)
		c.SetSession("id", id)
		c.Data["id"] = c.GetSession("id")
		//Sessionid=c.CruSession.SessionID()
		c.Redirect("/index", 302)
		//如果是学生用户
	} else if res1 == 0 && res2 == 1 {
		//fmt.Println(user.name,user.Password)
		c.SetSession("id", id)
		fmt.Println("id的值为：", id)

		//Sessionid=c.CruSession.SessionID()
		c.Redirect("/studentMainPage", 302)

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

func (c *MainController) Logout() {
	c.DelSession("id")
	c.Redirect("/", 302)
}

//func (c *MainController)SessionTest(){
//	if c.Data["id"]==nil{
//		c.Redirect("/login",302)
//	}
//}
