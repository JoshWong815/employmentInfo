package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

type Employment struct {
	Id   int
	Sid int
	Employed  string `orm:"size(128)"`
	Cname string  `orm:"size(128)"`
	Oname string   `orm:"size(128)"`
	Operation string `orm:"size(128)"`
}
var db MysqlDB

func init() {
	orm.RegisterModel(new(Employment))
	ok:=db.Connect()
	if ok!=true{
		log.Fatal("数据库连接失败")
		fmt.Println("数据库连接失败")
	}
}

//获得所有的签约信息
func GetAllEmployments() ([]*Employment,error){
	sql:="select employment.id,sid,operation,employed,company.name cname,offer.name oname from employment,company,offer where employment.cid=company.id and employment.oid=offer.id"

	list,err:=db.Query1(sql)
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	var employments []*Employment
	fmt.Println("list的值为：",list)
	for i:=range list{
		var employment Employment
        map1:=list[i]
        employment.Id,_=strconv.Atoi(string(map1["id"].([]uint8)))
        employment.Sid,_=strconv.Atoi(string(map1["sid"].([]uint8)))
		employment.Operation=string(map1["operation"].([]uint8))
        employment.Employed=string(map1["employed"].([]uint8))
        employment.Cname=string(map1["cname"].([]uint8))
        employment.Oname=string(map1["oname"].([]uint8))
        employments=append(employments,&employment)
	}
	return employments,nil

}

//获得所有的签约信息2
func GetAllEmployments2() ([]*Employment,error){
	sql:="select employment.id,sid,employed,company.name cname,offer.name oname from employment,company,offer where employment.cid=company.id and employment.oid=offer.id"

	list,err:=db.MysqlConnector.Query(sql)
	if err!=nil{
		fmt.Println("err:",err)
	}

	fmt.Println("list的值为：",list)

	return nil,nil

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
	sql:="delete from employment where id="+strconv.Itoa(id)
	fmt.Println(sql)
	res,err:=db.Exec(sql)
	fmt.Println(res)
	return nil
}