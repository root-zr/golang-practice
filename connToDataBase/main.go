package main

import (
	"fmt"
	"strings"

	"connToDatabase/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "question"
)

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	db, err := gorm.Open(mysql.Open(path), &gorm.Config{})

	if err != nil {
		fmt.Println("failed to connnect success")
	}

	db.AutoMigrate(&models.User{})
	fmt.Println("connnect success")

}

func main() {

	InitDB()
}
