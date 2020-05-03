package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"reflect"

	//"github.com/jmoiron/sqlx"
)
type MysqlDB struct{
	BaseDB
	MysqlConnector *sql.DB
}

func (db *MysqlDB) Connect() bool{
	db.Host="127.0.0.1"
	db.Port="3306"
	db.UserName="root"
	db.UserPwd="daomei,815"
	db.DBName="employmentinfo"
	var err error
	//fmt.Println("我在这里测试",db.UserName+":"+db.UserPwd+"@tcp("+db.Host+":"+db.Port+")/"+db.DBName)
	//db.Mysql,err=sql.Open("mysql",db.Name+":"+db.Pwd+"@tcp("+db.Host+":"+db.Port+")/"+db.DatabaseName)

	db.MysqlConnector,err=sql.Open("mysql",db.UserName+":"+db.UserPwd+"@tcp("+db.Host+":"+db.Port+")/"+db.DBName)

	if err!=nil{
		log.Fatal(err)
		return false
	}
	//fmt.Println(db)
	fmt.Println(db.UserName+":"+db.UserPwd+"@tcp("+db.Host+db.Port+")/"+db.DBName)
	return true
}
func (db *MysqlDB) Close(){
	db.Close()
}

func (db *MysqlDB) Query(sql string,args ...interface{}) (*sql.Rows,error){
	rows,err:=db.MysqlConnector.Query(sql,args...)
	return rows,err
}
func (db *MysqlDB) Query1(sql string,args ...interface{}) ([]map[string]interface{},error){
	rows,err:=db.MysqlConnector.Query(sql,args...)
	list:=rowsToMap(rows)
	return list,err
}
func (db *MysqlDB) Exec(sql string,args ...interface{})(sql.Result,error){
	result,err:=db.MysqlConnector.Exec(sql,args...)
	if err!=nil{
		logs.Error(err)
	}
	return result,err
}



//func rowsToMap(rows *sql.Rows) []interface{}{
//	columns,_:=rows.Columns()
//	coulmnsLength:=len(columns)
//	fmt.Println("columns的值为：",columns)
//	cache:=make([]interface{},coulmnsLength)
//	fmt.Println("cache的值为：",cache)
//	for i:=range cache{
//		var p interface{}
//		cache[i]=&p
//	}
//
//	for rows.Next(){
//		err:=rows.Scan(cache...)
//		if err!=nil{
//			logs.Error(err)
//		}
//		fmt.Println("cache的值为：",cache)
//		for i:=range cache{
//			a:=cache[i]
//			fmt.Println(a)
//		}
//
//	}
//	return nil
//
//}

func rowsToMap(rows *sql.Rows) []map[string]interface{}{
	columns,_:=rows.Columns()
	coulmnsLength:=len(columns)
	cache:=make([]interface{},coulmnsLength)
	fmt.Println("columns的值为：",columns)
	for index,_:=range cache{
		var a interface{}
		cache[index]=&a
	}
	var list []map[string]interface{}
	for rows.Next(){
		_ =rows.Scan(cache...)
		fmt.Println("cache的值为：",cache)
		item:=make(map[string]interface{})
		for i,data:=range cache{
			fmt.Println("i:",i,"cache[",i,"]:",data)
			fmt.Println("data的类型为：",reflect.TypeOf(data))
			fmt.Println(*data.(*interface{}))
			item[columns[i]]=*data.(*interface{})

		}
		list=append(list,item)
	}
	_=rows.Close()
	return list
}




//func rowsToMap(rows *sql.Rows) []map[string]interface{}{
//	columns,_:=rows.Columns()
//	coulmnsLength:=len(columns)
//	cache:=make([]interface{},coulmnsLength)
//	fmt.Println("columns的值为：",columns)
//	for index,_:=range cache{
//		var a interface{}
//		cache[index]=&a
//	}
//	var list []map[string]interface{}
//	for rows.Next(){
//		_ =rows.Scan(cache...)
//		item:=make(map[string]interface{})
//		for i,data:=range cache{
//			item[columns[i]]=*data.(*interface{})
//			fmt.Println(*data.(*interface{}))
//		}
//
//		list=append(list,item)
//	}
//	_=rows.Close()
//	return list
//}