package controllers

import "C"
import (
	"employmentInfo/models"
	"fmt"
)

//  singleStudentController operations for singleStudent
type SingleStudentController struct {
	//beego.Controller
	MainController

}

// URLMapping ...
func (c *SingleStudentController) URLMapping() {

c.Mapping("StudentMainPage",c.StudentMainPage)
c.Mapping("SingleStudentInfo",c.SingleStudentInfo)
c.Mapping("SingleStudentUpdating",c.SingleStudentUpdating)
}
func (c *SingleStudentController) SingleStudentUpdating(){
	Id:=c.GetSession("id")
	fmt.Println("Id的值：",Id)
	//intid, _ := strconv.Atoi(Id)
	//u := models.Student{Id: int64(intid)}
	u := models.Student{Id: Id.(string)}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println("parseForm这里的错误为：",err)
		c.Redirect("/singleStudentInfo" , 302)
	}
	fmt.Println(u)
	if err := models.UpdateStudentById(&u); err == nil {
		c.Redirect("/studentMainPage", 302)
	}else{
		fmt.Println(err)
		c.Redirect("/singleStudentInfo", 302)
	}
	//c.TplName="students.html"
}
//学生端的个人信息
func (c *SingleStudentController) SingleStudentInfo(){
	c.TplName="singleStudent/singleStudentInfo.html"
	c.SessionTest()
	id:=c.GetSession("id")
	fmt.Println("该名学生的id的值为：",id)
	//var student models.Student
	student,err:=models.GetStudentById(id.(string))
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(student)
	c.Data["student"]=student
}
//学生端的主页
func (c *SingleStudentController) StudentMainPage(){
	c.TplName="singleStudent/studentMainPage.html"
	c.SessionTest()
}

