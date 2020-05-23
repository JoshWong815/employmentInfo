package controllers

import "C"
import (
	"employmentInfo/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

//  OfferController operations for Offer
type OfferController struct {
	beego.Controller
}

// URLMapping ...
func (c *OfferController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAllOffers", c.GetAllOffers)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ShowOffers", c.ShowOffers)
	c.Mapping("UpdateOffer", c.UpdateOffer)
	c.Mapping("OfferUpdating", c.OfferUpdating)
	c.Mapping("DeleteOffer", c.DeleteOffer)
	c.Mapping("AddOffer", c.AddOffer)
	c.Mapping("OfferAdding", c.OfferAdding)
	c.Mapping("GetAllCompanyInOffer", c.GetAllCompanyInOffer)

}
func (c *OfferController) GetAllCompanyInOffer() {
	companys, _ := models.GetAllCompanyInOffer()
	c.Data["json"] = companys
	c.ServeJSON()
}

func (c *OfferController) OfferAdding() {
	var e models.Offer
	if err := c.ParseForm(&e); err != nil {
		fmt.Println("转换model失败")
		fmt.Println(err)
	}
	fmt.Println(e)
	Cid, _ := models.GetCidByCname(e.Cname)
	fmt.Println("Cid:", Cid)
	err := models.InsertAnOffer(e, Cid)
	if err != nil {
		fmt.Println(err)
	} else {
		c.Redirect("/getAllOffers", 302)
	}
}
func (c *OfferController) AddOffer() {
	c.TplName = "offer_add.html"
}
func (c *OfferController) DeleteOffer() {
	id := c.GetString("id")
	intid, _ := strconv.Atoi(id)
	if err := models.DeleteOffer(intid); err == nil {

		c.Redirect("/getAllOffers", 302)
	} else {
		c.Ctx.WriteString("删除失败！")
		c.Ctx.WriteString("id:" + id)

	}
}
func (c *OfferController) OfferUpdating() {
	Id := c.GetString("Id")
	fmt.Println("Id的值：", Id)
	intid, _ := strconv.Atoi(Id)
	//u := models.Offer{Id: int64(intid)}
	u := models.Offer{Id: intid}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println("parse的错误为：", err)
		c.Redirect("/updateOffer?id="+Id, 302)
	}
	fmt.Println(u)
	if err := models.UpdateOfferById(&u); err == nil {
		c.Redirect("/getAllOffers", 302)
	} else {
		c.Redirect("/updateOffer?id="+Id, 302)
	}
	c.TplName = "offers.html"
}
func (c *OfferController) ShowOffers() {
	c.TplName = "offers.html"
}
func (c *OfferController) UpdateOffer() {
	id := c.GetString("id")
	//id := c.Ctx.Input.Param(":id")
	fmt.Println("id:", id)
	intid, _ := strconv.Atoi(id)
	Offer, err := models.GetOfferById(intid)
	//Offer,err:=models.GetOfferById(id)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("该名学生的信息：",Offer)
	c.Data["list"] = Offer
	c.TplName = "offer_update.html"
}

// GetOne ...
// @Title Get One
// @Description get Offer by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Offer
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OfferController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOfferById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Offer
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Offer
// @Failure 403
// @router / [get]
func (c *OfferController) GetAllOffers() {
	c.Data["id"] = c.GetSession("id")
	offers, err := models.GetAllOffers()
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = offers
	}
	c.TplName = "offers.html"

}

// Put ...
// @Title Put
// @Description update the Offer
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Offer	true		"body for Offer content"
// @Success 200 {object} models.Offer
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OfferController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Offer{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateOfferById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Offer
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OfferController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOffer(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
