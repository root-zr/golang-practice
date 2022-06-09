package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        // gorm 要继承 Model， Model 里包括Id，创建时间，更新时间，删除时间等基本字段
	Username   string `gorm:"unique"`
	Password   string
}
