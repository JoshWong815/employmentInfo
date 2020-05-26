package controllers

import "C"
import (
	"employmentInfo/models"
	"fmt"
)

type ApprovalController struct {
	MainController
}

func (c *ApprovalController) URLMapping() {

	c.Mapping("GetAllApprovals", c.GetAllApprovals)
	c.Mapping("ShowAllApprovals", c.ShowAllApprovals)
	c.Mapping("ApprovalThis", c.ApprovalThis)
	c.Mapping("ApprovalsCanBeDone", c.ApprovalsCanBeDone)


}

//查询可操作性的审批记录
func (c *ApprovalController) ApprovalsCanBeDone(){
	approvals,_:=models.ApprovalsCanBeDone()
	c.Data["json"]=approvals
	c.ServeJSON()
}
//管理员端的通过审核
func (c *ApprovalController) ApprovalThis(){
	id:=c.GetString("id")
	fmt.Println("正在审核的审核id为：",id)
	//查询出该条审批记录
	a:=models.GetAnApproval(id)
	cid,_:=models.GetCidByCname(a.Cname)
	oid,_:=models.GetOidByOname(a.Oname,cid)
	//创建一个employment结构体，为之后的InsertAnEmployment函数做参数准备
	var e models.Employment
	e.Sid=a.Sid
	e.Operation=a.Operation
	e.Employed=a.Employed
	e.Cname=a.Cname
	e.Oname=a.Oname
	e.Reason=a.Reason
	err:=models.InsertAnEmployment(e,cid,oid)
	if err!=nil{
		fmt.Println("往employment表中插入数据时出错")
	}
	//将审核表的该条记录的mark值设为“否”，既已经通过申请
	_=models.SetApprovalMark(id)
	c.Redirect("showAllApprovals",302)
}

//管理员端的审批页面展示
func (c *ApprovalController) ShowAllApprovals(){

	c.Data["id"]=c.GetSession("id")
	c.SessionTest()
	c.TplName="approvals.html"
}

//管理员端的审批页面数据获取
func (c *ApprovalController) GetAllApprovals(){


	approvals, err := models.GetAllApprovals()
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = approvals
	}
	//c.TplName = "approvals.html"
	c.ServeJSON()

}
