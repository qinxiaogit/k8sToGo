package db

import (
	"fmt"
	"github.com/qinxiaogit/k8sToGo/binlog/conf"
	"gorm.io/driver/mysql" // gorm mysql 驱动包
	"gorm.io/gorm"         // gorm
)
var db *gorm.DB
func init(){
	// MySQL 配置信息
	username := conf.GetChildConf("mysql_user")             // 账号
	password := "xxxxxx" // 密码
	host := "127.0.0.1"             // 地址
	port := 3306                    // 端口
	DBname := "gorm1"               // 数据库名称
	timeout := "10s"                // 连接超时，10秒
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, DBname, timeout)
	// Open 连接
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect mysql.")
	}
}
