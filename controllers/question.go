package controllers

import "C"
import (
	"employmentInfo/models"
	"encoding/json"
	"fmt"
	"strconv"
)

//  QuestionController operations for Question
type QuestionController struct {
	//beego.Controller
	MainController
}

// URLMapping ...
func (c *QuestionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAllQuestions", c.GetAllQuestions)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ShowQuestions", c.ShowQuestions)
	c.Mapping("AnswerTheQuestion", c.AnswerTheQuestion)
	c.Mapping("QuestionAnswering", c.QuestionAnswering)
	c.Mapping("DeleteQuestion", c.DeleteQuestion)
	c.Mapping("AddQuestion", c.AddQuestion)
	c.Mapping("QuestionAdding", c.QuestionAdding)

}

func (c *QuestionController) QuestionAdding() {
	c.Data["id"] = c.GetSession("id")
	var s models.Question
	if err := c.ParseForm(&s); err != nil {
		fmt.Println("转换model失败")
		c.Ctx.WriteString("转换model失败")
		fmt.Println(err)
	}
	id, err := models.AddQuestion(&s)
	if err == nil && id > 0 {
		c.Redirect("/getAllQuestions", 302)
	} else if err != nil {
		fmt.Println("第二次err添加失败")
		//c.Ctx.WriteString("第二次err添加失败")
		fmt.Println(err)
	}
	c.Redirect("/getAllQuestions", 302)
}
func (c *QuestionController) AddQuestion() {
	c.Data["id"] = c.GetSession("id")
	c.Data["name"] = c.GetSession("name")
	c.SessionTest()
	c.TplName = "question_add.html"
}
func (c *QuestionController) DeleteQuestion() {
	id := c.GetString("id")
	//intid,_:=strconv.ParseInt(id,0,64)
	if err := models.DeleteQuestion(id); err == nil {

		c.Redirect("/getAllQuestions", 302)
	} else {
		c.Ctx.WriteString("删除失败！")
		fmt.Println(err)
		c.Ctx.WriteString("id:" + id)

	}
}
func (c *QuestionController) QuestionAnswering() {
	Id := c.GetString("Id")
	fmt.Println("Id的值：", Id)
	intid, _ := strconv.Atoi(Id)

	u := models.Question{Id: intid}
	aid := c.GetSession("id")

	fmt.Println("aid:", aid)
	if err := c.ParseForm(&u); err != nil {
		fmt.Println(err)
		c.Redirect("/updateQuestion?id="+Id, 302)
	}
	fmt.Println("u:", u)
	//aid, ok := aid.(string)
	//if ok==true{
	//	aid=aid.(string)
	//}else{
	//	aid=""
	//}
	aidStr := aid.(string)

	if err := models.UpdateQuestionById(&u, aidStr); err == nil {
		c.Redirect("/getAllQuestions", 302)
	} else {
		c.Redirect("/updateQuestion?id="+Id, 302)
	}
	c.TplName = "questions.html"
}
func (c *QuestionController) ShowQuestions() {
	c.Data["name"] = c.GetSession("name")
	c.Data["name"] = c.GetSession("name")
	c.TplName = "Questions.html"
}
func (c *QuestionController) AnswerTheQuestion() {
	id := c.GetString("id")
	//id := c.Ctx.Input.Param(":id")
	fmt.Println("id:", id)

	//intid, _ := strconv.Atoi(id)
	//Question,err:=models.GetQuestionById(int64(intid))
	Question, err := models.GetQuestionById(id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("该提问的信息：", Question)
	c.Data["id"] = c.GetSession("id")
	c.Data["name"] = c.GetSession("name")
	c.SessionTest()
	c.Data["list"] = Question
	c.TplName = "question_answer.html"
}

// Post ...
// @Title Post
// @Description create Question
// @Param	body		body 	models.Question	true		"body for Question content"
// @Success 201 {int} models.Question
// @Failure 403 body is empty
// @router / [post]
func (c *QuestionController) Post() {
	var v models.Question
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddQuestion(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Question by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Question
// @Failure 403 :id is empty
// @router /:id [get]
func (c *QuestionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetQuestionById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Question
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Question
// @Failure 403
// @router / [get]
func (c *QuestionController) GetAllQuestions() {
	c.Data["id"] = c.GetSession("id")
	c.Data["name"] = c.GetSession("name")
	c.SessionTest()
	sid := c.GetString("sid")
	questions, err := models.GetAllQuestions(sid)
	if err != nil {
		c.Data["json"] = err
	} else {
		c.Data["json"] = questions
	}
	c.TplName = "questions.html"

}
