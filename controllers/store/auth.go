package store

import "Jokaru-py/managingEntities/models"

type Auth interface {
	GetUser(login string) (*models.UsersDB, error)
	CreateUser(*models.UsersDB) error
}
