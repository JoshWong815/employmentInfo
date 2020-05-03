package models

type BaseDB struct {
	Host			string
	Port			string
	UserName		string
	UserPwd			string
	DBName			string
}
type BaseDBInterface interface {
	Connect() 	bool
	Close()		bool
	Query(sql string,args ...interface{})		[]map[string]interface{}
	Exec(sql string,args ...interface{})		bool
}
//func (db *BaseDB) Connect() bool{
//	return true
//}
//func (db *BaseDB) Close() bool{
//	return true
//}
//func (db *BaseDB) Query(sql string,connect *sql.DB,args ...interface{}) (*sql.Rows,error) {
//
//}
//func (db *BaseDB) Exec(sql string,connect *sql.DB,args ...interface{}) bool {
//	return true
//}

