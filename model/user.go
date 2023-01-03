package model

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:email`
	Password  string `json:password`
	Telephone string `json:telephone`
}

type UserModel struct {
	db *gorm.DB
}

func (model *UserModel) Insert(user User) (User, error) {

	err := model.db.Create(&user).Error
	if err != nil {
		log.Println("CREATE USER ERROR", err)
		return User{}, err
	}

	return user, nil
}
