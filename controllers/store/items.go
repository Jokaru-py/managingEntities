package store

import "Jokaru-py/managingEntities/models"

type Items interface {
	CreateObject(models.ObjectDB) error
}
