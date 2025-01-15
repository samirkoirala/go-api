package models

import (
	"github.com/Sulav-Adhikari/gouser/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string `json:"uname"`
	Password string `json:"password"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUser() *User {
	db.Create(&u)
	return u
}

func GetAllUsers() []User {
	var Users []User
	db.Find(&Users)
	return Users
}

func GetUserByUsername(username string) (*User, error) {
    var user User
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        if gorm.ErrRecordNotFound == err {
            return nil, nil 
        }
        return nil, err 
    }
    return &user, nil
}

