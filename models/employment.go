package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type Employment struct {
	Id        int
	Sid       int
	Operation string `orm:"size(128)"`
	Employed  string `orm:"size(128)"`
	Cname     string `orm:"size(128)"`
	Oname     string `orm:"size(128)"`
	Reason    string `orm:"size(128)"`
}

var db MysqlDB

func init() {
	orm.RegisterModel(new(Employment))
	ok := db.Connect()
	if ok != true {
		log.Fatal("数据库连接失败")
		fmt.Println("数据库连接失败")
	}
}

//获得所有的签约信息2
func GetAllEmployments2() ([]*Employment, error) {
	sql := "select employment.id,sid,operation,employed,company.name cname,offer.name oname,reason from employment,company,offer where employment.cid=company.id and employment.oid=offer.id"

	list, err := db.Query1(sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var employments []*Employment
	fmt.Println("list的值为：", list)
	for i := range list {
		var employment Employment
		map1 := list[i]
		employment.Id, _ = strconv.Atoi(string(map1["id"].([]uint8)))
		employment.Sid, _ = strconv.Atoi(string(map1["sid"].([]uint8)))
		employment.Operation = string(map1["operation"].([]uint8))
		employment.Employed = string(map1["employed"].([]uint8))
		employment.Cname = string(map1["cname"].([]uint8))
		employment.Oname = string(map1["oname"].([]uint8))

		employments = append(employments, &employment)
	}
	return employments, nil

}

//获得所有的签约信息
func GetAllEmployments() ([]*Employment, error) {
	var maps []orm.Params
	var employments []*Employment
	o := orm.NewOrm()
	sql := "select e.id id,sid,operation,employed,c.name cname,o.name oname,reason from (company c,offer o) right join employment e on e.cid=c.id and e.oid=o.id"
	a, err := o.Raw(sql).Values(&maps)
	fmt.Println("a:", a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(maps)
	for i := range maps {
		var employment Employment
		map1 := maps[i]
		employment.Id, _ = strconv.Atoi((map1["id"].(string)))
		employment.Sid, _ = strconv.Atoi((map1["sid"].(string)))
		//对操作类型进行断言
		operation, ok := map1["operation"].(string)
		if ok {
			employment.Operation = operation
		} else {
			employment.Operation = ""
		}

		employment.Employed = map1["employed"].(string)
		//对单位名称进行断言
		cname, ok := map1["cname"].(string)
		if ok {
			employment.Cname = cname
		} else {
			employment.Cname = ""
		}
		//对岗位名称进行断言
		oname, ok := map1["oname"].(string)
		if ok {
			employment.Oname = oname
		} else {
			employment.Oname = ""
		}
		//对解约原因进行断言
		reason, ok := map1["reason"].(string)
		if ok {
			employment.Reason = reason
		} else {
			employment.Reason = ""
		}

		employments = append(employments, &employment)
	}
	return employments, nil

}

//删除一项签约信息
func DeleteEmployment(id int) (err error) {
	//o := orm.NewOrm()
	//v := Employment{Id: id}
	//// ascertain id exists in the database
	//if err = o.Read(&v); err == nil {
	//	var num int64
	//	if num, err = o.Delete(&Employment{Id: id}); err == nil {
	//		fmt.Println("Number of records deleted in database:", num)
	//	}
	//}
	//return
	//sql:=fmt.Sprintf("delete from employment where id=",id)
	sql := "delete from employment where id=" + strconv.Itoa(id)
	fmt.Println(sql)
	res, err := db.Exec(sql)
	fmt.Println("删除employment时res的值为：", res)
	return nil
}

//修改一项签约信息
func UpdateEmployment(id int) {

}

//或得所有的单位id和名称
func GetAllCompanyIdAndNameInEmployment() ([]*Company, error) {
	var maps []orm.Params
	var companys []*Company
	o := orm.NewOrm()
	sql := "select id,name from company"
	n, err := o.Raw(sql).Values(&maps)
	fmt.Println("一共有", n, "条记录")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(maps)
	for i := range maps {
		var company Company
		map1 := maps[i]
		company.Id, _ = strconv.Atoi((map1["id"].(string)))
		company.Name = map1["name"].(string)
		companys = append(companys, &company)
	}
	return companys, nil
}

//获得所有offer的id和名称
func GetAllOfferIdAndNameInEmployment(s string) ([]*Offer, error) {
	var maps []orm.Params
	var offers []*Offer
	o := orm.NewOrm()
	//fmt.Print("s的值为：",s)
	sql := "select o.id id,o.name name from offer o,company c where o.cid=c.id and c.name='" + s + "'"
	fmt.Println(sql)
	n, err := o.Raw(sql).Values(&maps)
	fmt.Println("一共有", n, "条记录")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("maps:", maps)
	for i := range maps {
		var offer Offer
		map1 := maps[i]
		offer.Id, _ = strconv.Atoi((map1["id"].(string)))
		offer.Name = map1["name"].(string)
		offers = append(offers, &offer)
	}
	return offers, nil
}

//根据Cname获得Cid
func GetCidByCname(cname string) (int, error) {
	var cid int
	sql := "select id from company where name='" + cname + "'"
	o := orm.NewOrm()
	err := o.Raw(sql).QueryRow(&cid)
	if err != nil {
		fmt.Println("GetCidByCname出错", err)
		return 0, err
	}

	return cid, nil
}

//根据oname获得oid
func GetOidByOname(oname string, cid int) (int, error) {
	var oid int
	cidStr := strconv.Itoa(cid)
	sql := "select offer.id from offer where offer.name='" + oname + "' and cid='" + cidStr + "'"
	o := orm.NewOrm()
	err := o.Raw(sql).QueryRow(&oid)
	if err != nil {
		fmt.Println(err)
		return 0, nil
	}
	return oid, nil

}

//往数据库中插入一条employment记录
func InsertAnEmployment(e Employment, Cid, Oid int) error {
	sid := strconv.Itoa(e.Sid)
	cid := strconv.Itoa(Cid)
	oid := strconv.Itoa(Oid)
	var sql string
	if (Cid == 0) && (Oid == 0) {
		sql = "insert into employment(sid,operation,employed,cid,oid,reason,time) values(" + sid + ",'" + e.Operation + "','" + e.Employed + "'," + "null" + "," + "null" + "," + "'" + e.Reason + "',current_timestamp())"
	} else {
		sql = "insert into employment(sid,operation,employed,cid,oid,reason,time) values(" + sid + ",'" + e.Operation + "','" + e.Employed + "'," + cid + "," + oid + "," + "'" + e.Reason + "',current_timestamp())"
	}
	fmt.Println(sql)
	o := orm.NewOrm()
	res, err := o.Raw(sql).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
		fmt.Println("res:", res)
	} else {
		fmt.Print(err)
		return err
	}
	return nil
}

func CheckSid(sid string) (string, error) {
	var maps []orm.Params
	var maps2 []orm.Params
	var maps3 []orm.Params
	sql := "select * from employment where sid=" + sid + " and operation='签约' and employed='是'"
	sql2 := "select * from employment where sid=" + sid + " and operation='解约'"
	sql3 := "select * from employment where sid=" + sid + " and operation='从未签约'"
	fmt.Println(sql)
	fmt.Println(sql2)
	fmt.Println(sql3)
	o := orm.NewOrm()
	res, _ := o.Raw(sql).Values(&maps)
	res2, _ := o.Raw(sql2).Values(&maps2)
	res3, _ := o.Raw(sql3).Values(&maps3)
	fmt.Println("res的值为：", res)
	fmt.Println("res2的值为：", res2)
	fmt.Println("res3的值为：", res3)
	//1.从未签约过，可签约 2.已签约，只能解约 3.已解约，可签约

	//签约的记录总数减去解约的记录总数，如果等于一，说明该名学生当前已经签约了，只能进行解约操作，返回t
	if res-res2 == 1 {
		return "employed", nil
	}

	//签约的记录总数减去解约的记录总数，如果等于0，并且签约的次数不等于0，说明该名学生当前已经解约了，可以签约，返回f
	if res-res2 == 0 && res != 0 {
		return "unemployed", nil
	}
	//签约的记录总数减去解约的记录总数，如果等于0，并且从未签约过，说明该名学生从未签约过，可以签约，返回f
	if res-res2 == 0 && res3 == 0 {
		return "neverBothEmployOrNever", nil
	}
	//签约的记录总数减去解约的记录总数，如果等于0，并且从未签约过，说明该名学生从未签约过，只能进行签约
	if res-res2 == 0 && res3 == 1 {
		return "neverOnlyEmploy", nil
	}
	return "", nil
}

//当已签约的用户要解约时，查出他当前已签约的单位和岗位
func GetSidEmployment(sid string) (string, string, error) {
	var maps []orm.Params
	var employments []*Employment
	sql := "select c.name cname,o.name oname from employment e,company c,offer o where e.cid=c.id and e.oid=o.id and sid=" + sid + " ORDER BY e.id desc limit 1"
	fmt.Println(sql)
	o := orm.NewOrm()
	res, err := o.Raw(sql).Values(&maps)
	fmt.Println("res:", res)
	if err != nil {
		return "", "", err
	}
	for i := range maps {
		var e Employment
		map1 := maps[i]
		e.Cname = map1["cname"].(string)
		e.Oname = map1["oname"].(string)
		employments = append(employments, &e)
	}
	return employments[0].Cname, employments[0].Oname, nil
}

func GetEmploymentById(id string) (*Employment, error) {
	var e Employment

	sql := "select e.id id,sid,operation,employed,c.name cname,o.name oname,reason from employment e,company c, offer o where e.cid=c.id and e.oid= o.id and e.id=" + id
	fmt.Println(sql)
	o := orm.NewOrm()
	err := o.Raw(sql).QueryRow(&e)
	if err != nil {
		fmt.Println(err)
		return &e, err
	}
	fmt.Println("查询出来的这条记录为：", e)
	return &e, err

}

func UpdateEmploymentById(e *Employment) error {
	cid, err := GetCidByCname(e.Cname)
	if err != nil {
		fmt.Println(err)
	}
	cidStr := strconv.Itoa(cid)
	oid, err := GetOidByOname(e.Oname, cid)
	oidStr := strconv.Itoa(oid)
	idStr := strconv.Itoa(e.Id)
	if err != nil {
		fmt.Println(err)
	}
	sql := "update employment set cid=" + cidStr + ",oid=" + oidStr + " where id=" + idStr
	fmt.Println(sql)
	o := orm.NewOrm()
	res, err := o.Raw(sql).Exec()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("res的值为：", res)
	return err
}
