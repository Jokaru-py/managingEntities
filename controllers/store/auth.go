package store

import "Jokaru-py/managingEntities/models"

type Auth interface {
	GetUser(login string) (*models.Users, error)
	CreateUser(*models.Users) error
}
