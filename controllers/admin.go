package controllers

import "C"
import (
	"employmentInfo/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//  AdminController operations for Admin
type AdminController struct {
	//beego.Controller
	MainController
}

// URLMapping ...
func (c *AdminController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAllAdmins", c.GetAllAdmins)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)

	c.Mapping("UpdateAdmin", c.UpdateAdmin)
	c.Mapping("AdminUpdating", c.AdminUpdating)
	c.Mapping("DeleteAdmin", c.DeleteAdmin)
	c.Mapping("AddAdmin", c.AddAdmin)
	c.Mapping("AdminAdding", c.AdminAdding)

}

func (c *AdminController) AdminAdding() {
	c.Data["id"] = c.GetSession("id")
	var s models.Admin
	if err := c.ParseForm(&s); err != nil {
		fmt.Println("转换model失败")
		c.Ctx.WriteString("转换model失败")
		fmt.Println(err)
	}
	lastAdminId:=models.GetLastAdminId()
	s.Id=strconv.Itoa(lastAdminId+1)
	id, err := models.AddAdmin(&s)
	if err == nil && id > 0 {
		c.Redirect("/getAllAdmins", 302)
	} else if err != nil {
		fmt.Println("第二次err添加失败")
		//c.Ctx.WriteString("第二次err添加失败")
		fmt.Println(err)
	}
	c.Redirect("/getAllAdmins", 302)
}
func (c *AdminController) AddAdmin() {
	c.Data["id"] = c.GetSession("id")
	c.Data["name"] = c.GetSession("name")
	c.SessionTest()
	c.TplName = "admin_add.html"
}
func (c *AdminController) DeleteAdmin() {
	id := c.GetString("id")
	//intid,_:=strconv.ParseInt(id,0,64)
	if err := models.DeleteAdmin(id); err == nil {

		c.Redirect("/getAllAdmins", 302)
	} else {
		c.Ctx.WriteString("删除失败！")
		c.Ctx.WriteString("id:" + id)

	}
}
func (c *AdminController) AdminUpdating() {
	Id := c.GetString("Id")
	fmt.Println("Id的值：", Id)
	//intid, _ := strconv.Atoi(Id)
	//u := models.Admin{Id: int64(intid)}
	u := models.Admin{Id: Id}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println(err)
		c.Redirect("/updateAdmin?id="+Id, 302)
	}
	fmt.Println(u)
	if err := models.UpdateAdminById(&u); err == nil {
		c.Redirect("/getAllAdmins", 302)
	} else {
		c.Redirect("/updateAdmin?id="+Id, 302)
	}
	c.TplName = "admins.html"
}

func (c *AdminController) UpdateAdmin() {
	id := c.GetString("id")
	//id := c.Ctx.Input.Param(":id")
	fmt.Println("id:", id)
	fmt.Println("id:", id)
	//intid, _ := strconv.Atoi(id)
	//Admin,err:=models.GetAdminById(int64(intid))
	Admin, err := models.GetAdminById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("该名学生的信息：", Admin)
	c.Data["id"]=c.GetSession("id")
	c.Data["name"]=c.GetSession("name")
	c.Data["list"] = Admin
	c.SessionTest()
	c.TplName = "admin_update.html"
}

// Post ...
// @Title Post
// @Description create Admin
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 201 {int} models.Admin
// @Failure 403 body is empty
// @router / [post]
func (c *AdminController) Post() {
	var v models.Admin
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddAdmin(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Admin by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Admin
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AdminController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetAdminById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Admin
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Admin
// @Failure 403
// @router / [get]
func (c *AdminController) GetAllAdmins() {
	c.SessionTest()
	//fmt.Println("id:",c.Data["id"])
	//c.SessionTest()

	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 100
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

	l, err := models.GetAllAdmin(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	//c.ServeJSON()
	c.Data["name"]=c.GetSession("name")
	c.SessionTest()
	c.TplName = "admin.html"

}

// Put ...
// @Title Put
// @Description update the Admin
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Admin	true		"body for Admin content"
// @Success 200 {object} models.Admin
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AdminController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Admin{Id: idStr}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateAdminById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Admin
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AdminController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteAdmin(idStr); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
