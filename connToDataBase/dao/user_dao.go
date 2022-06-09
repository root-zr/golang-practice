package dao

import (
	"connToDatabase/datasource"
	"connToDatabase/models"
	"log"
)

type UserRepository interface {
	GetUserByUserNameAndPwd(username string, password string) (user models.User)
	GetUserByUsername(username string) (user models.User)
	Save(user models.User) (int, models.User)
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

type userRepository struct{}

//登录
func (n userRepository) GetUserByUserNameAndPwd(username string, password string) (user models.User) {
	db := datasource.GetDB()
	db.Where("username = ? and password = ?", username, password).First(&user)
	return
}
func (n userRepository) GetUserByUsername(username string) (user models.User) {
	db := datasource.GetDB()
	db.Where("username = ?", username).First(&user)
	return
}

//添加/修改
func (n userRepository) Save(user models.User) (int, models.User) {
	code := 0

	tx := datasource.GetDB().Begin()
	if user.ID != 0 {
		var oldUser models.User
		datasource.GetDB().First(&oldUser, user.ID)
		user.CreatedAt = oldUser.CreatedAt
		user.Username = oldUser.Username
		if user.Username == "" {
			user.Username = oldUser.Username
		}
		if user.Password == "" {
			user.Password = oldUser.Password
		}
	}

	if err := tx.Save(&user).Error; err != nil {
		log.Println(err)
		code = -1
	}
	return code, user
}
