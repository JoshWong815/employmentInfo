package models
import (
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"log"
	"os"
)
type Config struct {
	Db            struct {
		Host        string            `json:"host"`
		Port        string            `json:"port"`
		User        string            `json:"user"`
		Password    string            `json:"password"`
		Name        string            `json:"name"`
		Prefix      string            `json:"prefix"`
		MaxOpenConn int               `json:"max_open_conn"`
		Params      map[string]string `json:"params"`
		Debug       bool              `json:"debug"`
	} `json:"db"`
}

var instanceDb *gorm.DB
var DefaultConfig *Config
func NewDb()(*gorm.DB, error){
	if instanceDb==nil{
		config:=NewConfig("test.json")
		fmt.Println(config)
		mysqlConfig := mysql.NewConfig()
		mysqlConfig.User = config.Db.User
		mysqlConfig.DBName = config.Db.Name
		mysqlConfig.Passwd = config.Db.Password
		mysqlConfig.Params = config.Db.Params
		mysqlConfig.Net = "tcp"
		mysqlConfig.Addr = config.Db.Host + ":" + config.Db.Port
		db, err := gorm.Open("mysql", mysqlConfig.FormatDSN())
		db.DB().SetMaxOpenConns(config.Db.MaxOpenConn)
		db.SingularTable(true)
		instanceDb = db
		return instanceDb, err
	}
	return instanceDb, nil
}
func NewConfig(file string) *Config {
	if DefaultConfig==nil{
		configText,err:=ioutil.ReadFile(file)
		if err!=nil{
			log.Println("配置文件读取错误,启动默认配置:", err.Error())
			os.Exit(1)
		}
		err=json.Unmarshal(configText,&DefaultConfig)
		return DefaultConfig
	}
	return DefaultConfig
}