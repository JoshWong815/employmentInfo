package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Question struct {
	Id       int    `orm:"size(128);pk"`
	Question string `orm:"size(128)"`
	Answer   string `orm:"size(128)"`
	Sid      string `orm:"size(128)"`
	Aid      string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Question))
}

// AddQuestion insert a new Question into database and returns
// last inserted Id on success.
func AddQuestion(m *Question) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetQuestionById retrieves Question by Id. Returns error if
// Id doesn't exist
func GetQuestionById(id string) (v *Question, err error) {
	var q *Question
	o := orm.NewOrm()
	sql := "select * from question where id=" + id
	fmt.Println(sql)
	err = o.Raw(sql).QueryRow(&q)
	if err != nil {
		fmt.Println("查询question时出错：", err)
		return q, err
	}
	return q, err

}

//获得所有的答疑记录
func GetAllQuestions(sid string) ([]*Question, error) {
	var maps []orm.Params
	var Questions []*Question
	o := orm.NewOrm()
	var sql string
	if sid == "" {
		sql = "select * from question"
	} else {
		sql = "select * from question where sid=" + sid
	}

	a, err := o.Raw(sql).Values(&maps)
	fmt.Println("a:", a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(maps)
	for i := range maps {
		var Question Question
		map1 := maps[i]
		Question.Id, _ = strconv.Atoi((map1["id"].(string)))

		//对学生的问题进行断言
		question, ok := map1["question"].(string)
		if ok {
			Question.Question = question
		} else {
			Question.Question = ""
		}

		//对管理员的回复进行断言
		answer, ok := map1["answer"].(string)
		if ok {
			Question.Answer = answer
		} else {
			Question.Answer = ""
		}
		//对学生id进行断言
		sid, ok := map1["sid"].(string)
		if ok {
			Question.Sid = sid
		} else {
			Question.Sid = ""
		}
		//对管理员id进行断言
		aid, ok := map1["aid"].(string)
		if ok {
			Question.Aid = aid
		} else {
			Question.Aid = ""
		}

		Questions = append(Questions, &Question)
	}
	return Questions, nil

}

// UpdateQuestion updates Question by Id and returns error if
// the record to be updated doesn't exist
func UpdateQuestionById(m *Question, aid string) (err error) {
	o := orm.NewOrm()

	idStr := strconv.Itoa(m.Id)
	sql := "update question set answer='" + m.Answer + "',aid=" + aid + " where id=" + idStr
	fmt.Println(sql)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		fmt.Println(err)
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("res:", rowsAffected)
	return err
}

// DeleteQuestion deletes Question by Id and returns error if
// the record to be deleted doesn't exist
func DeleteQuestion(id string) (err error) {
	o := orm.NewOrm()
	//idInt,_:=strconv.Atoi(id)
	sql := "delete from question where id=" + id
	fmt.Println(sql)
	if res, err := o.Raw(sql).Exec(); err == nil {
		fmt.Println("res:", res)
		return nil
	} else {
		return err
	}

}
