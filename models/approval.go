package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Approval struct {
	Id        int
	Sid       int
	Operation string `orm:"size(128)"`
	Employed  string `orm:"size(128)"`
	Cname     string `orm:"size(128)"`
	Oname     string `orm:"size(128)"`
	Reason    string `orm:"size(128)"`
	Time      string `orm:"size(128)"`
	Mark 	  string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Approval))
}
func ApprovalsCanBeDone() ([]*Approval,error){
	sql := "select a.id id,sid,operation,employed,c.name cname,o.name oname,reason,time,mark from (company c,offer o) right join approval a on a.cid=c.id and a.oid=o.id where mark='是'"
	o := orm.NewOrm()
	var maps []orm.Params
	var approvals  []*Approval
	n, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println("获取可操作的审批记录时出错：", err)
		return approvals,err
	}
	fmt.Println("n:",n)
	for i := range maps {
		var approval Approval
		map1 := maps[i]
		approval.Id, _ = strconv.Atoi((map1["id"].(string)))
		approval.Sid, _ = strconv.Atoi((map1["sid"].(string)))
		//对操作类型进行断言
		operation, ok := map1["operation"].(string)
		if ok {
			approval.Operation = operation
		} else {
			approval.Operation = ""
		}

		approval.Employed = map1["employed"].(string)
		//对单位名称进行断言
		cname, ok := map1["cname"].(string)
		if ok {
			approval.Cname = cname
		} else {
			approval.Cname = ""
		}
		//对岗位名称进行断言
		oname, ok := map1["oname"].(string)
		if ok {
			approval.Oname = oname
		} else {
			approval.Oname = ""
		}
		//对解约原因进行断言
		reason, ok := map1["reason"].(string)
		if ok {
			approval.Reason = reason
		} else {
			approval.Reason = ""
		}

		//对时间进行断言
		time, ok := map1["time"].(string)
		if ok {
			approval.Time = time
		} else {
			approval.Time = ""
		}

		//对是否是最新记录进行断言
		mark, ok := map1["mark"].(string)
		if ok {
			approval.Mark = mark
		} else {
			approval.Mark = ""
		}

		approvals = append(approvals, &approval)
	}
	return approvals,nil
}


//获得所有的审批信息
func GetAllApprovals() ([]*Approval, error) {
	var maps []orm.Params
	var Approvals []*Approval
	o := orm.NewOrm()
	sql :="select a.id id,sid,operation,employed,c.name cname,o.name oname,reason,time,mark from (company c,offer o) right join approval a on a.cid=c.id and a.oid=o.id"
	fmt.Println(sql)
	a, err := o.Raw(sql).Values(&maps)
	fmt.Println("a:", a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(maps)
	for i := range maps {
		var approval Approval
		map1 := maps[i]
		approval.Id, _ = strconv.Atoi((map1["id"].(string)))
		approval.Sid, _ = strconv.Atoi((map1["sid"].(string)))
		//对操作类型进行断言
		operation, ok := map1["operation"].(string)
		if ok {
			approval.Operation = operation
		} else {
			approval.Operation = ""
		}

		approval.Employed = map1["employed"].(string)
		//对单位名称进行断言
		cname, ok := map1["cname"].(string)
		if ok {
			approval.Cname = cname
		} else {
			approval.Cname = ""
		}
		//对岗位名称进行断言
		oname, ok := map1["oname"].(string)
		if ok {
			approval.Oname = oname
		} else {
			approval.Oname = ""
		}
		//对解约原因进行断言
		reason, ok := map1["reason"].(string)
		if ok {
			approval.Reason = reason
		} else {
			approval.Reason = ""
		}

		//对操作时间进行断言
		time, ok := map1["time"].(string)
		if ok {
			approval.Time = time
		} else {
			approval.Time = ""
		}

		//对是否是最新记录进行断言
		mark, ok := map1["mark"].(string)
		if ok {
			approval.Mark = mark
		} else {
			approval.Mark = ""
		}

		Approvals = append(Approvals, &approval)
	}
	return Approvals, nil

}

//添加一条审批信息（这信息是从学生发过来的）
func InsertAnApproval(e Employment, Cid, Oid int) error{
	sid := strconv.Itoa(e.Sid)
	cid := strconv.Itoa(Cid)
	oid := strconv.Itoa(Oid)
	//先执行sql0的目的是把审批表中该名学生当前最新的一条记录的mark值设为否
	sql0:="update approval set mark='否' where sid="+sid+" and id=(select id from(select id from approval a where sid="+sid+" ORDER BY id desc limit 1) a)"
	fmt.Println(sql0)
	//
	sql1 := "insert into approval(sid,operation,employed,cid,oid,reason,time,mark) values("+sid+",'"+e.Operation+"','"+e.Employed+"',"+cid+","+oid+",'"+e.Reason+"',current_timestamp(),'是')"
	fmt.Println(sql1)
	o := orm.NewOrm()
	res0, err0 := o.Raw(sql0).Exec()
	res1, err1 := o.Raw(sql1).Exec()
	if err0 == nil {
		num, _ := res0.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		fmt.Println("res0:", res0)
	} else {
		fmt.Print(err0)
		return err0
	}
	if err1 == nil {
		num, _ := res1.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		fmt.Println("res:", res1)
	} else {
		fmt.Print(err1)
		return err1
	}
	return nil
}
//获得一条审批信息
func GetAnApproval(id string) *Approval{
	fmt.Println("要查询的审批id为：",id)

	sql:="select a.id,sid,operation,employed,c.name cname,o.name oname,reason,mark from approval a,company c,offer o where a.cid=c.id and a.oid=o.id and a.id="+id
	fmt.Println(sql)
	var a *Approval
	o:=orm.NewOrm()
	err:=o.Raw(sql).QueryRow(&a)
	fmt.Println("查询id为",id,"的审批记录时出错：",err)
	fmt.Println("查询出来的审批记录为：",a)
	return  a
}

//将审核表中的某条记录的mark值设为“否”，既已经通过申请
func SetApprovalMark(id string) error{
	fmt.Println("要修改mark值的approval的记录id为：",id)
	sql:="update approval set mark='否' where id="+id
	fmt.Println(sql)
	o:=orm.NewOrm()
	res,err:=o.Raw(sql).Exec()
	if err!=nil{
		fmt.Println("修改mark值时出错：",err)
		return err
	}
	resNum,_:=res.RowsAffected()
	fmt.Println("成功修改了",resNum,"条记录")
	return nil
	}







