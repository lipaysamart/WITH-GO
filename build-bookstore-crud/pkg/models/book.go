package models

import (
	"gorm.io/gorm"
)

type Book struct {
	// Model 里的结构体， 其中包括字段 id, CreateAt, UpdateAt, DeletedAt
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
