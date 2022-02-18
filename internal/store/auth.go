package store

import "gorm.io/gorm"

type ConnStore struct {
	db *gorm.DB
}

func NewConnStore(db *gorm.DB) *ConnStore {
	return &ConnStore{
		db: db,
	}
}
