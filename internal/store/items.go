package store

import (
	"Jokaru-py/managingEntities/models"
	"errors"

	"gorm.io/gorm"
)

// Создать объект в БД
func (cs *ConnStore) CreateObject(object *models.Object) (*models.Object, error) {
	err := cs.db.Create(object).Error
	if err != nil {
		return nil, err
	}

	return object, nil
}

// Получить объект в БД
func (cs *ConnStore) GetObjectByName(object *models.Object) (*models.Object, error) {
	var obj models.Object
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
func (cs *ConnStore) GetObjectByID(object *models.Object) (*models.Object, error) {
	var obj models.Object
	err := cs.db.Where("id = ? and user_id = ?", object.ID, object.UserID).Take(&obj).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, nil
	}

	return &obj, nil
}

// Получить объект в БД по ID
func (cs *ConnStore) DeleteObjectByID(object *models.Object) error {
	return cs.db.Delete(object).Where("id = ", object.ID).Error
}

// Получить все объекты
func (cs *ConnStore) GetAllObjectByID(params *models.Object) ([]*models.Object, error) {
	var res []*models.Object
	err := cs.db.Where("user_id = ?", params.UserID).Find(&res).Error
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Поменять владельца объекта
func (cs *ConnStore) UpdateObject(params *models.Object) error {
	return cs.db.Where("id = ?", params.ID).Updates(params).Error
}
