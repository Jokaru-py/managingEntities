package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Login  string   `gorm:"not null;column:login"` // Логин
	Pass   string   `gorm:"not null;column:pass"`  // Пароль
	Object []Object `gorm:"foreignKey:UserID"`
}

type Object struct {
	gorm.Model
	Name   string `gorm:"not null;column:name"` // Название объекта
	UserID uint
}
