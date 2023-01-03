package model

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"id"`
	Name      string `json:"name" gorm:"type:varchar(50)"`
	Email     string `json:"email" gorm:"type:varchar(50); not null"`
	Password  string `json:"password" gorm:"type:varchar(255); not null"`
	Telephone string `json:"telephone" gorm:"type:varchar(15)"`
	IsActive  bool   `gorm:"type:bool; default:true"`
}

type UserModel struct {
	DB *gorm.DB
}

func (model *UserModel) Insert(user User) (User, error) {
	tx := model.DB.Create(&user)
	if tx.Error != nil {
		log.Println("CREATE USER QUERY ERROR", tx.Error)
		return User{}, tx.Error
	}

	return user, nil
}

func (model *UserModel) Save(user User) (User, error) {
	tx := model.DB.Save(&user)
	if tx.Error != nil {
		log.Println("UPDATE USER QUERY ERROR", tx.Error)
		return User{}, tx.Error
	}

	return user, nil
}

func (model *UserModel) Delete(id int) error {
	tx := model.DB.Delete(id)
	if tx.Error != nil {
		log.Println("DELETE USER QUERY ERROR", tx.Error)
		return tx.Error
	}

	return nil
}

func (model *UserModel) FindById(id int) (User, error) {
	user := User{}
	tx := model.DB.Find(&user, id)
	if tx.Error != nil {
		log.Println("FIND BY ID USER QUERY ERROR", tx.Error)
		return User{}, tx.Error
	}

	return user, nil
}

func (model *UserModel) FindAll() ([]User, error) {
	users := []User{}
	tx := model.DB.Find(&users)
	if tx.Error != nil {
		log.Println("FIND ALL USER QUERY ERROR", tx.Error)
		return []User{}, tx.Error
	}

	return users, nil
}
