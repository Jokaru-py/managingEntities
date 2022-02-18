package store

import "gorm.io/gorm"

type ConnStore struct {
	db *gorm.DB
}

func (as *ConnStore) GetHeadsDepAgency(id uint) (*[]string, error) {

	return nil, nil
}
