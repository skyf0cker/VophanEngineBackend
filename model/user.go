package model

import (
	"errors"
	"github.com/jinzhu/gorm"
	"regexp"
)

type User struct {
	gorm.Model
	Username   string       `json:"username"`
	Password   string       `json:"password"`
	Email      string       `json:"email"`
	UserChoice []UserChoice `gorm:"ForeignKey:UserRefer"json:"user_choice"`
}

//暂时不校验邮箱
func CreateUser(u *User) error {
	if CheckUsername(u.Username) {
		var rows int
		db.Where("username = ?", u.Username).Find(&[]User{}).Count(&rows)
		if rows > 0 {
			return errors.New("账号已存在")
		} else {
			if CheckPassword(u.Password) {
				return db.Save(u).Error
			} else {
				return errors.New("密码不合法")
			}
		}
	} else {
		return errors.New("账号不合法")
	}
}

func CheckPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}

func CheckUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", username); !ok {
		return false
	}
	return true
}

func GetUserByName(username string) (user *User, err error) {
	user = &User{}
	err = db.Where("username = ?", username).First(user).Error
	return
}

func DeleteUserByName(username string) error {
	return db.Where("username = ?", username).Delete(&User{}).Error
}

func UpdateUser(username string, change map[string]interface{}) error {
	return db.Model(&User{}).Where("username = ?", username).Update(change).Error
}
