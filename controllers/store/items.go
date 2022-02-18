package store

import "Jokaru-py/managingEntities/models"

type Items interface {
	GetObjectByName(*models.ObjectDB) error
	GetObjectByID(*models.ObjectDB) error
	DeleteObjectByID(*models.ObjectDB) error
	GetAllObjectByID(*models.ObjectDB) ([]*models.ObjectDB, error)
}
