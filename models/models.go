package models

/**
获取数据库配置并连接数据库
*/

import (
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/imanner/gin-xinshulaila/pkg/setting"
)

// 定义db对象
var db *gorm.DB

// 定义表的公共部分
type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// 初始化
func init() {
	// 定义变量 引入配置信息
	var (
		err error
		dbType, dbName, user, password, host, tablePrefix string
	)

	// 配置文件获取数据库名
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// 获取数据库类型
	dbType = sec.Key("TYPE").String()
	// 数据库名
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	// gorm 链接数据库
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	// gorm 拼接完整的数据库
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return tablePrefix + defaultTableName;
	}

	// 单例打开水数据库
	db.SingularTable(true)
	// 最大连接数
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// 关闭连接
func CloseDB() {
	defer db.Close()
}
