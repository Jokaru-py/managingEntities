package store

import "Jokaru-py/managingEntities/models"

type Items interface {
	GetObjectByName(*models.Object) error
	GetObjectByID(*models.Object) error
	DeleteObjectByID(*models.Object) error
	GetAllObjectByID(*models.Object) ([]*models.Object, error)
}
