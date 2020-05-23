package controllers

import "C"
import (
	"employmentInfo/models"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//  SkillController operations for Skill
type SkillController struct {
	beego.Controller
}

// URLMapping ...
func (c *SkillController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAllSkills", c.GetAllSkills)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ShowSkills", c.ShowSkills)
	c.Mapping("UpdateSkill", c.UpdateSkill)
	c.Mapping("SkillUpdating", c.SkillUpdating)
	c.Mapping("DeleteSkill", c.DeleteSkill)
	c.Mapping("AddSkill", c.AddSkill)
	c.Mapping("SkillAdding", c.SkillAdding)

}

func (c *SkillController) SkillAdding() {
	var s models.Skill
	if err := c.ParseForm(&s); err != nil {
		fmt.Println("转换model失败")
		c.Ctx.WriteString("转换model失败")
		fmt.Println(err)
	}
	id, err := models.AddSkill(&s)
	if err == nil && id > 0 {
		c.Redirect("/getAllSkills", 302)
	} else if err != nil {
		fmt.Println("第二次err添加失败")
		//c.Ctx.WriteString("第二次err添加失败")
		fmt.Println(err)
	}
	c.Redirect("/getAllSkills", 302)
}
func (c *SkillController) AddSkill() {
	c.TplName = "skill_add.html"
}
func (c *SkillController) DeleteSkill() {
	id := c.GetString("id")
	intid, _ := strconv.ParseInt(id, 0, 64)
	if err := models.DeleteSkill(intid); err == nil {

		c.Redirect("/getAllSkills", 302)
	} else {
		c.Ctx.WriteString("删除失败！")
		c.Ctx.WriteString("id:" + id)

	}
}
func (c *SkillController) SkillUpdating() {
	Id := c.GetString("Id")
	fmt.Println("Id的值：", Id)
	intid, _ := strconv.Atoi(Id)
	//u := models.Skill{Id: int64(intid)}
	u := models.Skill{Id: int64(intid)}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println(err)
		c.Redirect("/updateSkill?id="+Id, 302)
	}
	fmt.Println(u)
	if err := models.UpdateSkillById(&u); err == nil {
		c.Redirect("/getAllSkills", 302)
	} else {
		c.Redirect("/updateSkill?id="+Id, 302)
	}
	c.TplName = "skills.html"
}
func (c *SkillController) ShowSkills() {
	c.TplName = "skills.html"
}
func (c *SkillController) UpdateSkill() {
	id := c.GetString("id")
	//id := c.Ctx.Input.Param(":id")
	fmt.Println("id:", id)
	intid, _ := strconv.Atoi(id)
	Skill, err := models.GetSkillById(int64(intid))
	//Skill,err:=models.GetSkillById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("该单位的信息：", Skill)
	c.Data["list"] = Skill
	c.TplName = "skill_update.html"
}

// Post ...
// @Title Post
// @Description create Skill
// @Param	body		body 	models.Skill	true		"body for Skill content"
// @Success 201 {int} models.Skill
// @Failure 403 body is empty
// @router / [post]
func (c *SkillController) Post() {
	var v models.Skill
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddSkill(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Skill by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Skill
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SkillController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetSkillById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Skill
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Skill
// @Failure 403
// @router / [get]
func (c *SkillController) GetAllSkills() {
	c.Data["id"] = c.GetSession("id")
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

	l, err := models.GetAllSkill(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	//c.ServeJSON()
	c.TplName = "skills.html"

}

// Put ...
// @Title Put
// @Description update the Skill
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Skill	true		"body for Skill content"
// @Success 200 {object} models.Skill
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SkillController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Skill{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateSkillById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Skill
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SkillController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteSkill(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
