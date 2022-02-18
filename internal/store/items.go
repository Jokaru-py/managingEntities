package store

import "Jokaru-py/managingEntities/models"

// Создать пользователя в БД
func (cs *ConnStore) CreateObject(object *models.ObjectDB) error {
	return cs.db.Create(object).Error
}
