package controllers

import (

	"github.com/astaxie/beego"
	"employmentInfo/models"
	"strconv"
)
type EmploymentController struct {
beego.Controller
}
func (c *EmploymentController) URLMapping() {
	c.Mapping("GetAllEmployments", c.GetAllEmployments)
	c.Mapping("DeleteEmployment", c.DeleteEmployment)
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


