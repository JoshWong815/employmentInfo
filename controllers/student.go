package controllers

import "C"
import (
	"crypto/md5"
	"employmentInfo/models"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

//  StudentController operations for Student
type StudentController struct {
	//beego.Controller
	MainController

}

// URLMapping ...
func (c *StudentController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAllStudents", c.GetAllStudents)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ShowStudents", c.ShowStudents)
	c.Mapping("UpdateStudent",c.UpdateStudent)
	c.Mapping("StudentUpdating",c.StudentUpdating)
	c.Mapping("DeleteStudent",c.DeleteStudent)
	c.Mapping("AddStudent",c.AddStudent)
	c.Mapping("StudentAdding",c.StudentAdding)
	c.Mapping("FileUpload",c.FileUpload)

}

func (c *StudentController) FileUpload(){

	f, h, _ := c.GetFile("myfile")//获取上传的文件
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求
	var AllowExtMap map[string]bool = map[string]bool{
		".jpg":true,
		".jpeg":true,
		".png":true,
	}
	if _,ok:=AllowExtMap[ext];!ok{
		c.Ctx.WriteString( "后缀名不符合上传要求" )
		return
	}
	//创建目录
	uploadDir := "static/upload/" + time.Now().Format("2006/01/02/")
	err := os.MkdirAll( uploadDir , 777)
	if err != nil {
		c.Ctx.WriteString( fmt.Sprintf("%v",err) )
		return
	}
	//构造文件名称
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000 )
	hashName := md5.Sum( []byte( time.Now().Format("2006_01_02_15_04_05_") + randNum ) )

	fileName := fmt.Sprintf("%x",hashName) + ext
	//c.Ctx.WriteString(  fileName )

	fpath := uploadDir + fileName
	defer f.Close()//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveToFile("myfile", fpath)
	if err != nil {
		c.Ctx.WriteString( fmt.Sprintf("%v",err) )
	}
	//c.Ctx.WriteString( "上传成功~！！！！！！！" )
	fmt.Println("上传成功")
	c.Redirect("/getAllStudents",302)
	c.TplName="students.html"
}

func (c *StudentController) StudentAdding(){
	c.Data["id"]=c.GetSession("id")
	var s models.Student
	if err := c.ParseForm(&s); err != nil {
		fmt.Println("转换model失败")
		c.Ctx.WriteString("转换model失败")
		fmt.Println(err)
	}
	id, err := models.AddStudent(&s)
	if err == nil && id > 0 {
		c.Redirect("/getAllStudents", 302)
	} else if err!=nil{
		fmt.Println("第二次err添加失败")
		//c.Ctx.WriteString("第二次err添加失败")
		fmt.Println(err)
	}
	c.Redirect("/getAllStudents",302)
}
func (c *StudentController) AddStudent(){
	c.Data["id"]=c.GetSession("id")
	c.TplName="student_add.html"
}
func (c *StudentController) DeleteStudent(){
	id:=c.GetString("id")
	//intid,_:=strconv.ParseInt(id,0,64)
	if err := models.DeleteStudent(id); err == nil {

		c.Redirect("/getAllStudents",302)
	} else {
		c.Ctx.WriteString("删除失败！")
		c.Ctx.WriteString("id:"+id)

	}
}
func (c *StudentController) StudentUpdating() {
	Id := c.GetString("Id")
	fmt.Println("Id的值：",Id)
	//intid, _ := strconv.Atoi(Id)
	//u := models.Student{Id: int64(intid)}
	u := models.Student{Id: Id}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println(err)
		c.Redirect("/updateStudent?id="+Id , 302)
	}
	fmt.Println(u)
	if err := models.UpdateStudentById(&u); err == nil {
		c.Redirect("/getAllStudents", 302)
	}else{
		c.Redirect("/updateStudent?id="+Id , 302)
	}
		c.TplName="students.html"
}
func (c *StudentController) ShowStudents() {
	c.TplName="students.html"
}
func (c *StudentController) UpdateStudent(){
	id:=c.GetString("id")
	//id := c.Ctx.Input.Param(":id")
	fmt.Println("id:",id)
	fmt.Println("id:",id)
	//intid, _ := strconv.Atoi(id)
	//student,err:=models.GetStudentById(int64(intid))
	student,err:=models.GetStudentById(id)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("该名学生的信息：",student)
    c.Data["list"]=student
    c.TplName="student_update.html"
}
// Post ...
// @Title Post
// @Description create Student
// @Param	body		body 	models.Student	true		"body for Student content"
// @Success 201 {int} models.Student
// @Failure 403 body is empty
// @router / [post]
func (c *StudentController) Post() {
	var v models.Student
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddStudent(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Student by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Student
// @Failure 403 :id is empty
// @router /:id [get]
func (c *StudentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetStudentById(idStr)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Student
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Student
// @Failure 403
// @router / [get]
func (c *StudentController) GetAllStudents() {
	c.SessionTest()
	//fmt.Println("id:",c.Data["id"])
	//c.SessionTest()
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

	l, err := models.GetAllStudent(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}

	//c.ServeJSON()
	c.TplName="students.html"

}

// Put ...
// @Title Put
// @Description update the Student
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Student	true		"body for Student content"
// @Success 200 {object} models.Student
// @Failure 403 :id is not int
// @router /:id [put]
func (c *StudentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Student{Id: idStr}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateStudentById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Student
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *StudentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	//id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteStudent(idStr); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
