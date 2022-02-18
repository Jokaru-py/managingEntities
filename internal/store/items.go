package store

import (
	"Jokaru-py/managingEntities/models"
	"errors"

	"gorm.io/gorm"
)

// Создать объект в БД
func (cs *ConnStore) CreateObject(object *models.ObjectDB) error {
	return cs.db.Create(object).Error
}

// Получить объект в БД
func (cs *ConnStore) GetObjectByName(object *models.ObjectDB) (*models.ObjectDB, error) {
	var obj models.ObjectDB
	err := cs.db.Where("name = ?", object.Name).Take(&obj).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &obj, nil
}

// Получить объект в БД по ID
func (cs *ConnStore) GetObjectByID(object *models.ObjectDB) (*models.ObjectDB, error) {
	var obj models.ObjectDB
	err := cs.db.Where("id = ?", object.ID).Take(&obj).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &obj, nil
}

// Получить объект в БД по ID
func (cs *ConnStore) DeleteObjectByID(object *models.ObjectDB) error {
	return cs.db.Delete(object).Where("id = ", object.ID).Error
}

// Получить все объекты
func (cs *ConnStore) GetAllObjectByID(*models.ObjectDB) ([]*models.ObjectDB, error) {
	var res []*models.ObjectDB
	err := cs.db.Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}
