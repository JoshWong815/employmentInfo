package controllers

import "C"
import (
	"employmentInfo/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
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
c.Mapping("singleStudentCompanys",c.SingleStudentCompanys)
c.Mapping("SingleStudentOffers",c.SingleStudentOffers)
c.Mapping("StudentEmployTheOffer",c.StudentEmployTheOffer)
c.Mapping("SingleStudentEmploymentAdding",c.SingleStudentEmploymentAdding)
c.Mapping("GetSidEmployed",c.GetSidEmployed)
}

//检查该名学生是否签约了
func (c *SingleStudentController) GetSidEmployed(){
	sid:=c.GetString("sid")
	var maps []orm.Params
	var arr []string
	fmt.Println("sid的值为：",sid)
	o:=orm.NewOrm()
	sql:="SELECT((SELECT count(*) from employment where  operation='签约' and sid="+sid+")-(SELECT count(*) from employment where  operation='解约' and sid="+sid+")) as num"
	fmt.Println(sql)
	res,err:=o.Raw(sql).Values(&maps)
	if err!=nil{
		fmt.Println(err)
	}
	for i:=range maps{

		map1:=maps[i]
		a:=map1["num"].(string)
		fmt.Println("a的值为：",a)
		arr=append(arr,a)
	}
	fmt.Println("arr:",arr)

	fmt.Println("res的值为：",res)
	c.Data["json"]=arr[0]
	c.ServeJSON()
}
//学生自己签约某个岗位的页面跳转
func (c *SingleStudentController) StudentEmployTheOffer(){
	c.SessionTest()
	c.TplName="singleStudent/singleStudentEmployTheOffer.html"
}
//学生自己签约某个岗位的具体实现
func (c *SingleStudentController) SingleStudentEmploymentAdding(){
	var e models.Employment
	e.Reason=c.GetString("Reason")
	if err := c.ParseForm(&e); err != nil {
		fmt.Println("转换model失败")
		fmt.Println(err)
	}
	fmt.Println(e)
	Cid,_:=models.GetCidByCname(e.Cname)
	Oid,_:=models.GetOidByOname(e.Oname,Cid)
	fmt.Println("Cid:",Cid,"Oid:",Oid)
	err:=models.InsertAnEmployment(e,Cid,Oid)
	if err!=nil{
		return
	}else {

		c.Redirect("/studentMainPage", 302)
	}
}


func (c *SingleStudentController) SingleStudentOffers(){
	c.Data["id"]=c.GetSession("id")
	offers,err:=models.GetAllOffers()
	if err!=nil{
		c.Data["json"]=err
	}else{
		c.Data["json"]=offers
	}
	c.TplName="singleStudent/singleStudentOffers.html"
}

func (c *SingleStudentController) SingleStudentCompanys(){
	c.Data["id"]=c.GetSession("id")
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllCompany(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.TplName="singleStudent/singleStudentCompanys.html"

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

