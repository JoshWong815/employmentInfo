package controllers

import "C"
import (
	"employmentInfo/models"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)
type EmploymentController struct {
beego.Controller
}
func (c *EmploymentController) URLMapping() {
	c.Mapping("GetAllEmployments", c.GetAllEmployments)
	c.Mapping("DeleteEmployment", c.DeleteEmployment)
	//c.Mapping("UpdateEmployment", c.UpdateEmployment)
	c.Mapping("AddEmployment", c.AddEmployment)
	c.Mapping("EmploymentAdding", c.EmploymentAdding)
	c.Mapping("GetNowCompany", c.GetNowCompany)
	c.Mapping("CheckSid", c.CheckSid)
	c.Mapping("GetSidEmployment", c.GetSidEmployment)
	c.Mapping("GetAllCompanyNames", c.GetAllCompanyNames)
	c.Mapping("UpdateEmployment", c.UpdateEmployment)


}
//修改签约信息页面跳转
func (c *EmploymentController) UpdateEmployment(){
	id:=c.GetString("id")
	fmt.Println("id:",id)
	intid, _ := strconv.Atoi(id)
	Employment,err:=models.GetEmploymentById(intid)

	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("该单位的信息：",Company)
	c.Data["list"]=Company
	c.TplName="company_update.html"
	c.TplName="employment_update.html"
}

//当已签约的用户要解约时，查出他当前已签约的单位和岗位
func (c *EmploymentController) GetSidEmployment(){
	sid:=c.GetString("sid")
	cname,oname,_:=models.GetSidEmployment(sid)
	fmt.Println(sid+"已签约"+cname+"公司的"+oname+"岗位")
	type CandSname struct {
	Cname string
	Oname string
	}
	var cs CandSname
	cs.Cname=cname
	cs.Oname=oname
	//cs:=&{cname,oname}
	fmt.Println("cs:",cs)
	c.Data["json"]=&cs
	c.ServeJSON()
}

func (c *EmploymentController) CheckSid(){
	sid:=c.GetString("sid")
	fmt.Println("sid的值为：",sid)
	str,err:=models.CheckSid(sid)
	if err!=nil{
		fmt.Println(err)
	}
	//if str=="t"{
		c.Data["json"]=str
		c.ServeJSON()
		fmt.Println("json:",c.Data["json"])
	//}
}

func (c *EmploymentController) EmploymentAdding(){
	var e models.Employment
	//e.Sid,_=strconv.Atoi(c.GetString("Sid"))
	//e.Operation=c.GetString("Operation")
	//e.Employed=c.GetString("Employed")
	//e.Reason=c.GetString("Reason")
	//e.Cname=c.GetString("Cname")
	//e.Oname=c.GetString("Oname")
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
		c.Redirect("/getAllEmployments", 302)
	}
}

//根据select下拉框选择的单位名称获得该单位所有的岗位名称
func (c *EmploymentController) GetNowCompany(){
	cname:=c.GetString("cname")
	fmt.Println("cname:",cname)
	c.TplName="employment_add.html"
	Offers:=c.giveCnameToFindOffer(cname)
	c.Data["json"]=Offers
	fmt.Println("offer的值为：",Offers)
	c.ServeJSON()
}
//中转函数，将单位名称传给那个函数已让其去查找该单位所有的岗位名称
func (c *EmploymentController)giveCnameToFindOffer(cname string) []*models.Offer{
	offers:=c.GetAllOfferIdAndNameInEmployment(cname)
	return offers
}
//在添加签约信息页面中获取当前单位所有的岗位名称
func (c *EmploymentController) GetAllOfferIdAndNameInEmployment(s string)  []*models.Offer{
	var offers []*models.Offer
	offers,err:=models.GetAllOfferIdAndNameInEmployment(s)
	if err!=nil{
		fmt.Println(err)
	}
	c.Data["offers"]=offers
	return offers
}

//在添加签约信息页面中获取当前所有的单位名称
func (c *EmploymentController) GetAllCompanyNames()  {
	var companys []*models.Company
	companys,err:=models.GetAllCompanyIdAndNameInEmployment()
	if err!=nil{
		fmt.Println(err)
	}
	c.Data["json"]=companys
	fmt.Println("company的值为:",companys)
	c.ServeJSON()
}

//添加签约信息页面的初始化
func (c *EmploymentController) AddEmployment(){
	c.TplName="employment_add.html"
	//c.GetAllCompanyIdAndNameInEmployment()

}



func (c *EmploymentController) GetAllEmployments() {
	c.Data["id"]=c.GetSession("id")
		employments,err:=models.GetAllEmployments()
		if err!=nil{
			c.Data["json"]=err
		}else{
			c.Data["json"]=employments
		}
		c.TplName="employments.html"


}

func (c *EmploymentController) DeleteEmployment(){
	id:=c.GetString("id")
	intid,_:=strconv.ParseInt(id,0,0)
	if err := models.DeleteEmployment(int(intid)); err == nil {

		c.Redirect("/getAllEmployments",302)
	} else {
		c.Ctx.WriteString("删除失败！")
		c.Ctx.WriteString("id:"+id)

	}
}




