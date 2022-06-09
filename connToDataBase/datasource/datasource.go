package datasource

import (
	"connToDatabase/conf"
	"connToDatabase/models"
	"fmt"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func createTable() {
	db.AutoMigrate(&models.User{})
}

func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{conf.Sysconfig.DBUserName, ":", conf.Sysconfig.DBPassword, "@tcp(",
		conf.Sysconfig.DBIp, ":", conf.Sysconfig.DBPort, ")/", conf.Sysconfig.DBName, "?charset=utf8"}, "")

	var err error
	db, err = gorm.Open(mysql.Open(path), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connnect success")
	}

	// createTable()
	fmt.Println("connnect success")

}
