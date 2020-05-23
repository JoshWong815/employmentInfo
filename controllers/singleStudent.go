package controllers

import "C"
import (
	"employmentInfo/models"
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
)

//  singleStudentController operations for singleStudent
type SingleStudentController struct {
	//beego.Controller
	MainController
}

// URLMapping ...
func (c *SingleStudentController) URLMapping() {

	c.Mapping("StudentMainPage", c.StudentMainPage)
	c.Mapping("SingleStudentInfo", c.SingleStudentInfo)
	c.Mapping("SingleStudentUpdating", c.SingleStudentUpdating)
	c.Mapping("singleStudentCompanys", c.SingleStudentCompanys)
	c.Mapping("SingleStudentOffers", c.SingleStudentOffers)
	c.Mapping("StudentEmployTheOffer", c.StudentEmployTheOffer)
	c.Mapping("SingleStudentEmploymentAdding", c.SingleStudentEmploymentAdding)
	c.Mapping("GetSidEmployed", c.GetSidEmployed)
	c.Mapping("SingleStudentSkills", c.SingleStudentSkills)
	c.Mapping("SingleStudentQuestions", c.SingleStudentQuestions)
	c.Mapping("SingleStudentAddQuestion", c.SingleStudentAddQuestion)
	c.Mapping("SingleStudentQuestionAdding", c.SingleStudentQuestionAdding)
	c.Mapping("GetThisStudentQuestion", c.GetThisStudentQuestion)
	c.Mapping("GetSkillsOfThisType", c.GetSkillsOfThisType)
	c.Mapping("ShowSingleStudentQuestions", c.ShowSingleStudentQuestions)
	c.Mapping("GetAllCompanysInOffer", c.GetAllCompanysInOffer)
	c.Mapping("ChooseTheCompanyInOffer", c.ChooseTheCompanyInOffer)
	c.Mapping("GetAllCitys", c.GetAllCitys)
	c.Mapping("GetThisCitysCompany", c.GetThisCitysCompany)

}

func (c *SingleStudentController) GetThisCitysCompany() {
	city := c.GetString("city")
	fmt.Println("city:", city)
	sql := "select * from company where city='" + city + "'"
	fmt.Println(sql)
	o := orm.NewOrm()
	var companys []*models.Company
	res, err := o.Raw(sql).QueryRows(&companys)
	if err != nil {
		fmt.Println("查询所在城市为", city, "的单位时出错:", err)
	}

	fmt.Println("res", res)
	fmt.Println("所在城市为", city, "的单位有:", companys)
	c.Data["json"] = companys
	c.ServeJSON()
}

func (c *SingleStudentController) GetAllCitys() {
	sql := "select distinct city from company"
	fmt.Println(sql)
	var res []orm.Params
	var citys []string

	o := orm.NewOrm()
	n, err := o.Raw(sql).Values(&res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("n:", n)
	for i := range res {
		map1 := res[i]
		citys = append(citys, map1["city"].(string))
	}
	fmt.Println("citys:", citys)
	c.Data["json"] = citys
	c.ServeJSON()
}

func (c *SingleStudentController) ChooseTheCompanyInOffer() {
	var offers []*models.Offer
	cname := c.GetString("cname")
	fmt.Println("cname:", cname)
	sql := "select o.id,o.name, c.name cname,o.note from offer o,company c where o.cid=c.id and c.name='" + cname + "'"
	fmt.Println(sql)
	o := orm.NewOrm()
	n, err := o.Raw(sql).QueryRows(&offers)
	if err != nil {
		fmt.Println("查询", cname, "公司的岗位时出错：", err)
	}
	fmt.Println("n:", n)
	//fmt.Println(offers)
	for i := range offers {
		fmt.Println(offers[i])
	}
	c.Data["json"] = offers
	c.ServeJSON()
}

func (c *SingleStudentController) GetAllCompanysInOffer() {
	c.Data["id"] = c.GetSession("id")
	c.SessionTest()

	var res []orm.Params
	var names []string
	sql := "select name from company"
	fmt.Println(sql)
	o := orm.NewOrm()
	n, err := o.Raw(sql).Values(&res)
	for i := range res {
		//fmt.Println(res[i])
		map1 := res[i]
		names = append(names, map1["name"].(string))
	}
	if err != nil {
		fmt.Println("查询所有单位的名字时出错：", err)
	}
	fmt.Println("n的值为：", n)
	fmt.Println("names的值为：", names)

	c.Data["json"] = names
	c.ServeJSON()

}

func (c *SingleStudentController) GetSkillsOfThisType() {
	var s []*models.Skill
	typeOfSkill := c.GetString("type")
	fmt.Println(typeOfSkill)
	sql := "select * from skill where type='" + typeOfSkill + "'"
	fmt.Println(sql)
	o := orm.NewOrm()
	res, err := o.Raw(sql).QueryRows(&s)
	if err != nil {
		fmt.Println("查询"+typeOfSkill+"类型的技能时出错：", err)
	}

	fmt.Println("res:", res)
	fmt.Println(&s)
	//遍历符合type条件的就业技能
	for i := range s {
		fmt.Println(s[i])
	}
	c.Data["json"] = s
	c.ServeJSON()
}

func (c *SingleStudentController) ShowSingleStudentQuestions() {
	c.Data["id"] = c.GetSession("id")
	c.TplName = "singleStudent/singleStudentQuestions.html"
}

func (c SingleStudentController) GetThisStudentQuestion() {
	//var maps []orm.Params
	var questions []*models.Question
	sid := c.GetString("sid")
	fmt.Println("sid的值为：", sid)
	sql := "select * from question where sid=" + sid
	fmt.Println(sql)
	o := orm.NewOrm()
	//res,err:=o.Raw(sql).Values(&maps)
	res, err := o.Raw(sql).QueryRows(&questions)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("res:", res)

	fmt.Println("questions:", questions)
	c.Data["json"] = questions
	c.ServeJSON()
}

func (c *SingleStudentController) SingleStudentQuestionAdding() {
	sid := c.GetSession("id")
	sidStr := sid.(string)

	fmt.Println("sid的值为:", sid)
	question := c.GetString("Question")
	sql := "insert into question(id,question,sid) values(null,'" + question + "'," + sidStr + ")"
	fmt.Println(sql)
	o := orm.NewOrm()
	res, err := o.Raw(sql).Exec()
	if err != nil {
		fmt.Println(err)
	}
	resNum, _ := res.RowsAffected()
	fmt.Println("res:", resNum)
	if resNum != 0 {
		c.Redirect("showSingleStudentQuestions", 302)
	}
}

func (c *SingleStudentController) SingleStudentAddQuestion() {
	c.Data["id"] = c.GetSession("id")
	c.SessionTest()
	c.TplName = "singleStudent/singleStudentAddQuestion.html"
}

func (c *SingleStudentController) SingleStudentQuestions() {
	c.Data["id"] = c.GetSession("id")
	c.SessionTest()
	sid := c.GetString("sid")
	questions, err := models.GetAllQuestions(sid)
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = questions
		c.ServeJSON()
	}
	c.TplName = "singleStudent/singleStudentQuestions.html"
}

func (c *SingleStudentController) SingleStudentSkills() {
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

	c.TplName = "singleStudent/singleStudentSkills.html"
}

//检查该名学生是否签约了
func (c *SingleStudentController) GetSidEmployed() {
	sid := c.GetString("sid")
	var maps []orm.Params
	var arr []string
	fmt.Println("sid的值为：", sid)
	o := orm.NewOrm()
	sql := "SELECT((SELECT count(*) from employment where  operation='签约' and sid=" + sid + ")-(SELECT count(*) from employment where  operation='解约' and sid=" + sid + ")) as num"
	fmt.Println(sql)
	res, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println(err)
	}
	for i := range maps {

		map1 := maps[i]
		a := map1["num"].(string)
		fmt.Println("a的值为：", a)
		arr = append(arr, a)
	}
	fmt.Println("arr:", arr)

	fmt.Println("res的值为：", res)
	c.Data["json"] = arr[0]
	c.ServeJSON()
}

//学生自己签约某个岗位的页面跳转
func (c *SingleStudentController) StudentEmployTheOffer() {
	c.SessionTest()
	c.TplName = "singleStudent/singleStudentEmployTheOffer.html"
}

//学生自己签约某个岗位的具体实现
func (c *SingleStudentController) SingleStudentEmploymentAdding() {
	var e models.Employment
	e.Reason = c.GetString("Reason")
	if err := c.ParseForm(&e); err != nil {
		fmt.Println("转换model失败")
		fmt.Println(err)
	}
	fmt.Println(e)
	Cid, _ := models.GetCidByCname(e.Cname)
	Oid, _ := models.GetOidByOname(e.Oname, Cid)
	fmt.Println("Cid:", Cid, "Oid:", Oid)
	err := models.InsertAnEmployment(e, Cid, Oid)
	if err != nil {
		return
	} else {

		c.Redirect("/studentMainPage", 302)
	}
}

func (c *SingleStudentController) SingleStudentOffers() {
	c.Data["id"] = c.GetSession("id")
	c.SessionTest()
	offers, err := models.GetAllOffers()
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = offers
	}
	c.Data["offer"] = offers

	c.TplName = "singleStudent/singleStudentOffers.html"
}

func (c *SingleStudentController) SingleStudentCompanys() {
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

	l, err := models.GetAllCompany(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.Data["Companys"] = l

	c.TplName = "singleStudent/singleStudentCompanys.html"

}

func (c *SingleStudentController) SingleStudentUpdating() {
	Id := c.GetSession("id")
	fmt.Println("Id的值：", Id)
	//intid, _ := strconv.Atoi(Id)
	//u := models.Student{Id: int64(intid)}
	u := models.Student{Id: Id.(string)}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println("parseForm这里的错误为：", err)
		c.Redirect("/singleStudentInfo", 302)
	}
	fmt.Println(u)
	if err := models.UpdateStudentById(&u); err == nil {
		c.Redirect("/studentMainPage", 302)
	} else {
		fmt.Println(err)
		c.Redirect("/singleStudentInfo", 302)
	}
	//c.TplName="students.html"
}

//学生端的个人信息
func (c *SingleStudentController) SingleStudentInfo() {
	c.TplName = "singleStudent/singleStudentInfo.html"
	c.SessionTest()
	id := c.GetSession("id")
	fmt.Println("该名学生的id的值为：", id)
	//var student models.Student
	student, err := models.GetStudentById(id.(string))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(student)
	c.Data["student"] = student
}

//学生端的主页
func (c *SingleStudentController) StudentMainPage() {
	c.TplName = "singleStudent/studentMainPage.html"
	c.SessionTest()
}
