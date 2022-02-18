package models

import "gorm.io/gorm"

type UsersDB struct {
	gorm.Model
	Login string `gorm:"not null;column:login"` // Логин
	Pass  string `gorm:"not null;column:pass"`  // Пароль
}

type ObjectDB struct {
	gorm.Model
	Name string `gorm:"not null;column:name"` // Название объекта
}
