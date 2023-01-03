package model

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id" gorm:"type:int; primary key; auto_increment"`
	Name      string `json:"name" form:"name"  gorm:"type:varchar(50); unique"`
	Email     string `json:"email" form:"email"  gorm:"type:varchar(50); not null"`
	Password  string `json:"password" form:"password" gorm:"type:varchar(255); not null"`
	Telephone string `json:"telephone" form:"telephone" gorm:"type:varchar(15)"`
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
	user := User{}
	tx := model.DB.Model(&user).Where("id = ?", id).Update("is_active", 0)
	if tx.Error != nil {
		log.Println("DELETE USER QUERY ERROR", tx.Error)
		return tx.Error
	}

	return nil
}

func (model *UserModel) FindById(id int) (User, error) {
	user := User{}
	tx := model.DB.Where("is_active = 1").Select("id", "name", "email", "telephone").Find(&user, id)
	if tx.Error != nil {
		log.Println("FIND BY ID USER QUERY ERROR", tx.Error)
		return User{}, tx.Error
	}

	return user, nil
}

func (model *UserModel) FindAll() ([]User, error) {
	users := []User{}
	tx := model.DB.Where("is_active = 1").Select("id", "name", "email", "telephone").Find(&users)
	if tx.Error != nil {
		log.Println("FIND ALL USER QUERY ERROR", tx.Error)
		return []User{}, tx.Error
	}

	return users, nil
}

func (model *UserModel) FindByEmail(email string) (User, error) {
	user := User{}
	tx := model.DB.Where("email = ?", email).Select("id", "name", "email", "password", "telephone").Find(&user)
	if tx.Error != nil {
		log.Println("FIND BY EMAIL USER QUERY ERROR", tx.Error)
		return User{}, tx.Error
	}

	return user, nil
}
