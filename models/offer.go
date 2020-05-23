package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Offer struct {
	Id    int
	Name  string `orm:"size(128)"`
	Cname string `orm:"size(128)"`
	Note  string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Offer))
}

//获得当前所有的单位名称
func GetAllCompanyInOffer() ([]*Company, error) {
	var maps []orm.Params
	var companys []*Company
	o := orm.NewOrm()
	sql := "select id,name from company"
	res, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println("查询所有company时出错：", err)
		return companys, err
	}
	fmt.Print("res的值为：", res)

	for i := range maps {
		var company Company
		map1 := maps[i]
		company.Id, _ = strconv.Atoi((map1["id"].(string)))
		//对岗位名称进行断言
		name, ok := map1["name"].(string)
		if ok {
			company.Name = name
		} else {
			company.Name = ""
		}
		fmt.Print(company)
		companys = append(companys, &company)
	}
	//fmt.Println("companys:",&companys)
	return companys, err

}

//往数据库中插入一条employment记录
func InsertAnOffer(e Offer, Cid int) error {

	cid := strconv.Itoa(Cid)

	var sql string
	sql = "insert into offer(name,cid,note) values('" + e.Name + "'," + cid + ",'" + e.Note + "')"
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

func GetOfferById(id int) (Offer, error) {

	var offer Offer
	o := orm.NewOrm()
	idStr := strconv.Itoa(id)
	sql := "select o.id,o.name,c.name cname,o.note from offer o,company c where o.cid=c.id and o.id= " + idStr
	fmt.Println(sql)
	err := o.Raw(sql).QueryRow(&offer)
	if err != nil {
		fmt.Println("查询单条offer是出错：", err)

	}
	fmt.Println("offer:", offer)
	return offer, err

}

// GetAllOffer retrieves all Offer matches certain condition. Returns empty list if
// no records exist
func GetAllOffers() ([]*Offer, error) {
	var maps []orm.Params
	var offers []*Offer
	o := orm.NewOrm()
	sql := "select o.id,o.name,c.name cname,o.note from company c,offer o where o.cid=c.id "
	a, err := o.Raw(sql).Values(&maps)
	fmt.Println("a:", a)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(maps)
	for i := range maps {
		var offer Offer
		map1 := maps[i]
		offer.Id, _ = strconv.Atoi((map1["id"].(string)))
		//对岗位名称进行断言
		name, ok := map1["name"].(string)
		if ok {
			offer.Name = name
		} else {
			offer.Name = ""
		}

		//对单位名称进行断言
		cname, ok := map1["cname"].(string)
		if ok {
			offer.Cname = cname
		} else {
			offer.Cname = ""
		}

		//对岗位概述进行断言
		note, ok := map1["note"].(string)
		if ok {
			offer.Note = note
		} else {
			offer.Note = ""
		}
		offers = append(offers, &offer)
	}
	return offers, nil

}

// UpdateOffer updates Offer by Id and returns error if
// the record to be updated doesn't exist
func UpdateOfferById(o *Offer) (err error) {
	orm := orm.NewOrm()
	idStr := strconv.Itoa(o.Id)
	cid, err := GetCidByCname(o.Cname)
	cidStr := strconv.Itoa(cid)
	if err != nil {
		fmt.Println(err)
		return err
	}
	sql := "update  offer set name='" + o.Name + "',cid=" + cidStr + ",note='" + o.Note + "' where id=" + idStr
	fmt.Println(sql)
	res, err := orm.Raw(sql).Exec()
	if err != nil {
		fmt.Println("err的值为：", err)
	}
	fmt.Println("res的值为：", res)

	return err
}

// DeleteOffer deletes Offer by Id and returns error if
// the record to be deleted doesn't exist
func DeleteOffer(id int) (err error) {
	o := orm.NewOrm()
	idStr := strconv.Itoa(id)
	sql := "delete from offer where id=" + idStr
	fmt.Println(sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(res)
	return err
}
