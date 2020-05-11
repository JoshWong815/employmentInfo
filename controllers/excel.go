package controllers

import (

	"fmt"
	"github.com/astaxie/beego"
	"log"
	"reflect"
	"strconv"
	"time"
	"github.com/xuri/excelize"
)
type ExcelController struct {
	beego.Controller
}

func (c *ExcelController) URLMapping() {
	c.Mapping("UploadExcel", c.UploadExcel)
}

type ExcelData interface {
	CreateMap(arr []string) map[string]interface{}
	ChangeTime(source string) time.Time
}
type ExcelStrcut struct {
	temp  [][]string
	Model interface{}
	Info  []map[string]string
}
type CacrmCaCrmUserHelper struct {
	JobCode       string
	Password    string
	Code string
	Name string
	TypeName string
	Type string
	BackCode string
}
var file *excelize.File

func (c *ExcelController)UploadExcel() {
	f, err := excelize.OpenFile("test.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	excel_map := f.GetSheetMap()
	for i, k := range excel_map {
		rows, err := f.GetRows(k)
		if err != nil {
			fmt.Println(err)
			return
		}
		e := ExcelStrcut{}
		temp := CacrmCaCrmUserHelper{}
		e.Model = temp
		e.ReadExcel(rows).CreateMap().SaveDb(i, &temp)
	}
}


func (excel *ExcelStrcut) ReadExcel(stringArray [][]string) *ExcelStrcut {
	excel.temp = stringArray
	return excel

}

func (excel *ExcelStrcut) CreateMap() *ExcelStrcut {
	//利用反射得到字段名
	for _, v := range excel.temp {
		var info = make(map[string]string)
		for i := 0; i < reflect.ValueOf(excel.Model).NumField(); i++ {

			obj := reflect.TypeOf(excel.Model).Field(i)
			//fmt.Printf("key:%s--val:%s\n",obj.Name,v[i])
			info[obj.Name] = v[i]
		}
		excel.Info = append(excel.Info, info)
	}
	return excel
}
func (excel *ExcelStrcut) ChangeTime(source string) time.Time {
	ChangeAfter, err := time.Parse("2006-01-02", source)
	if err != nil {
		log.Fatalf("转换时间错误:%s", err)
	}
	return ChangeAfter
}



	func(excel *ExcelStrcut) SaveDb(j int, temp CacrmCaCrmUserHelper) *ExcelStrcut{
		//忽略标题行
		for i := 1; i<len(excel.Info); i++{

		t := reflect.ValueOf(temp).Elem()
		for k, v := range excel.Info[i]{
		fmt.Println("666666666666666", excel.Info[i])
		//fmt.Println(t.FieldByName(k).t.FieldByName(k).Kind())
		//fmt.Println("key:%v---val:%v",t.FieldByName(k),t.FieldByName(k).Kind())

		switch t.FieldByName(k).Kind(){
	case reflect.String:
		t.FieldByName(k).Set(reflect.ValueOf(v))
	case reflect.Float64:
		tempV, err := strconv.ParseFloat(v, 64)
		if err != nil{
		log.Printf("string to float64 err：%v", err)
	}

		t.FieldByName(k).Set(reflect.ValueOf(tempV))
	case reflect.Uint64:
		reflect.ValueOf(v)
		tempV, err := strconv.ParseUint(v, 0, 64)
		if err != nil{
		log.Printf("string to uint64 err：%v", err)
	}
		t.FieldByName(k).Set(reflect.ValueOf(tempV))

	case reflect.Struct:
		tempV, err := time.Parse("2006-01-02", v)
		if err!=nil{
		log.Fatalf("string to time err:%v", err)
	}
		t.FieldByName(k).Set(reflect.ValueOf(tempV))
	default:
		fmt.Println(t.FieldByName(k).Kind())
	}
	}
		fmt.Println("5555555555555555")
		DB, ok := NewDb()
		if ok!=nil{
		fmt.Println("6666666", ok)
	}
		DB.SingularTable(true)
		if j<3{
		temp.Name = temp.Code
		temp.Code = ""
		temp.BackCode = temp.Type
		temp.Type = ""
	}
		err := DB.Debug().Create(&temp).Error
		if err != nil{
		log.Fatalf("save temp table err:%v", err)
	}
	}
		return excel
	}
