package model

import "github.com/jinzhu/gorm"

type UserChoice struct {
	gorm.Model
	Content string `json:"content"`
	UserRefer uint `json:"user_refer"`
}
