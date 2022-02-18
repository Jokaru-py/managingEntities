package store

import (
	"Jokaru-py/managingEntities/models"

	"gorm.io/gorm"
)

type ConnStore struct {
	db *gorm.DB
}

func NewConnStore(db *gorm.DB) *ConnStore {
	return &ConnStore{
		db: db,
	}
}

// Поиск user по логину
func (cs *ConnStore) GetUser(login string) (*models.UsersDB, error) {
	var res models.UsersDB
	err := cs.db.Where("login = ?", login).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Создать пользователя в БД
func (cs *ConnStore) CreateUser(user *models.UsersDB) error {
	return cs.db.Create(user).Error
}
